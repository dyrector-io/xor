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
	log.Info().Msgf("quiz select running %v", today)
	listAll := processor.MaskAndFilter(processor.ReadJSONData(), true, FilterIfStartLessThan)
	var indices []int
	if state.AppConfig.Method == RandomGenMethod {
		indices = PickRandom(QuestionCount, len(listAll)-1)
	} else {
		indices = GetPicksIfPresent(state.DBConn, today)
		if len(indices) == 0 {
			log.Info().Msg("generating new quiz")

			indices = PickByDate(today, QuestionCount, len(listAll)-1, database.GetExclusionList(state.DBConn))

			err := database.PersistPicks(state.DBConn, today, indices)
			if err != nil {
				log.Error().Err(err).Msg("persisting quiz picks for the day")
			}
		}
	}

	selected := processor.CNCFSequence{}

	for _, i := range indices {
		selected = append(selected, listAll[i])
	}

	state.QuizList = selected
}
