package game

import (
	"crypto/rand"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/dyrector-io/xor/api/pkg/processor"
	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
	"golang.org/x/exp/slices"
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

func PickByDate(amount, upperLimit int, excluded []int) []int {
	today := time.Now()

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

func Mask(original, mask string) string {
	return strings.ReplaceAll(original, mask, "...")
}

func QuizListResponse(list processor.CNCFSequence) []render.Renderer {
	result := []render.Renderer{}
	for _, item := range list {
		result = append(result, item)
	}
	return result
}

func GetQuiz(w http.ResponseWriter, r *http.Request) {
	listAll, err := processor.ReadJSONData()
	if err != nil {
		log.Error().Err(err).Send()
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	qNum := 5
	indices := PickByDate(qNum, len(listAll)-1, []int{})

	selected := processor.CNCFSequence{}

	for _, i := range indices {
		selected = append(selected, listAll[i])
	}

	err = render.RenderList(w, r, QuizListResponse(selected))
	if err != nil {
		log.Error().Err(err)
	}
}
