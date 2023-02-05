package game

import (
	"crypto/rand"
	"math/big"
	"time"

	"github.com/dyrector-io/xor/api/internal/config"
	"github.com/dyrector-io/xor/api/pkg/database"
	"github.com/dyrector-io/xor/api/pkg/processor"
	"github.com/rs/zerolog/log"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
)

const (
	FilterIfStartLessThan = 1000
	QuestionCount         = 5
)

const (
	RandomGenMethod = "RANDOM"
	LinearGenMethod = "LINEAR"
)

func PickRandom(amount, upperLimit int) []int {
	picked := []int{}
	min := 0
	bg := big.NewInt(int64(upperLimit) - int64(min))
	for i := 0; i < amount; i++ {
		gen, err := rand.Int(rand.Reader, bg)
		if err != nil {
			log.Error().Err(err).Send()
		}
		if i > 0 {
			if slices.Contains(picked, int(gen.Int64())) {
				i--
				continue
			}
		}
		picked = append(picked, int(gen.Int64()))
	}

	return picked
}

func sdbmHash(data []byte) uint64 {
	var hash uint64

	for _, b := range data {
		hash = uint64(b) + (hash << 6) + (hash << 16) - hash
	}

	return hash
}

func PickByDate(today time.Time, amount, upperLimit int, excluded []int) []int {
	if len(excluded)+amount >= upperLimit {
		log.Info().Msgf("ran out of item indices %v/%v", len(excluded), upperLimit)
		return []int{}
	}

	hash := sdbmHash([]byte(today.Format("2006-01-02")))
	slice := hash / uint64(amount)
	log.Info().Msgf("%d", hash)

	picked := []int{}
	for i := 0; i < amount; i++ {
		gen := hash % uint64(upperLimit)
		hash -= slice
		if i > 0 {
			for j := 1; slices.Contains(picked, int(gen)) || slices.Contains(excluded, int(gen)); j++ {
				gen = (hash + uint64(j)) % uint64(upperLimit)
			}
		}
		picked = append(picked, int(gen))
	}

	log.Info().Msgf("%v", picked)
	return picked
}

func GetPicksIfPresent(db *gorm.DB, today time.Time) []int {
	return database.GetPicksForDay(db, today)
}

func SelectAQuiz(state *config.AppState) {
	today := time.Now()
	log.Info().Msgf("SelectAQuiz running %v", today)
	listAll := processor.ReadJSONData()
	var indices []int
	if state.AppConfig.EndDate != "" {
		end, err := time.Parse("2006-01-02", state.AppConfig.EndDate)
		if err != nil {
			log.Err(err).Msg("invalid ending date format, expected: 2006-01-02")
		}
		if today.After(end) {
			log.Info().Msgf("End date passed %v>%v", today, state.AppConfig.EndDate)
			state.QuizList = processor.QuizSequence{}
			state.Ended = true
			return
		}
	}
	if state.QuizCounter*QuestionCount >= len(listAll) {
		state.QuizList = processor.QuizSequence{}
		state.Ended = true
		return
	} else if state.AppConfig.Method == RandomGenMethod {
		indices = PickRandom(QuestionCount, len(listAll)-1)
		log.Info().Msg("generating new random quiz")
	} else if state.AppConfig.Method == LinearGenMethod {
		start := state.QuizCounter * QuestionCount
		for i := start; i < start+QuestionCount; i++ {
			indices = append(indices, i)
		}
		state.QuizCounter++
	} else {
		indices = GetPicksIfPresent(state.DBConn, today)
		if len(indices) == 0 {
			log.Info().Msg("generating new quiz")

			indices = PickByDate(today, QuestionCount, len(listAll)-1, database.GetExclusionList(state.DBConn))

			err := database.PersistPicks(state.DBConn, today, indices)
			if err != nil {
				log.Error().Err(err).Msg("persisting quiz picks for the day")
			}
		} else {
			log.Info().Msg("found already generated quiz for today")
		}
	}

	if len(indices) == 0 {
		state.Ended = true
		log.Info().Msg("no quiz indices were picked, assuming the end")
		return
	}

	selected := processor.QuizSequence{}

	for _, i := range indices {
		selected = append(selected, listAll[i])
	}

	state.QuizList = selected
}
