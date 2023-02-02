package server

import (
	"net/http"
	"time"

	"github.com/dyrector-io/xor/api/internal/config"
	"github.com/dyrector-io/xor/api/internal/ctx"
	"github.com/dyrector-io/xor/api/pkg/database"
	"github.com/dyrector-io/xor/api/pkg/processor"
	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
)

type QuizListResponse struct {
	List processor.QuizSequence
	Date string
}

func (c *QuizListResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func QuizHistoryResponse(list database.HistoryList) []render.Renderer {
	result := []render.Renderer{}
	for _, item := range list {
		result = append(result, item)
	}
	return result
}

func GetQuiz(w http.ResponseWriter, r *http.Request) {
	appState := ctx.GetContextVar[*config.AppState](r.Context(), ctx.StateKey)
	if appState.Ended {
		http.Error(w, http.StatusText(http.StatusGone), http.StatusGone)
		return
	}
	err := render.Render(w, r, &QuizListResponse{
		List: appState.QuizList,
		Date: time.Now().Add(time.Hour * 24 * time.Duration(appState.QuizCounter)).Format(database.SimpleDateFormat),
	})
	if err != nil {
		log.Error().Err(err)
	}
}

func GetHistory(w http.ResponseWriter, r *http.Request) {
	appState := ctx.GetContextVar[*config.AppState](r.Context(), ctx.StateKey)
	err := render.RenderList(w, r, QuizHistoryResponse(database.GetHistoryDB(appState.DBConn)))
	if err != nil {
		log.Error().Err(err)
	}
}
