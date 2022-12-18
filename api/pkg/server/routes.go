package server

import (
	"net/http"

	"github.com/dyrector-io/xor/api/internal/config"
	"github.com/dyrector-io/xor/api/internal/ctx"
	"github.com/dyrector-io/xor/api/pkg/database"
	"github.com/dyrector-io/xor/api/pkg/processor"
	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
)

func QuizListResponse(list processor.CNCFSequence) []render.Renderer {
	result := []render.Renderer{}
	for _, item := range list {
		result = append(result, item)
	}
	return result
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
	err := render.RenderList(w, r, QuizListResponse(appState.QuizList))
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
