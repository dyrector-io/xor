package server

import (
	"net/http"

	"github.com/dyrector-io/xor/api/internal/config"
	"github.com/dyrector-io/xor/api/internal/ctx"
	"github.com/dyrector-io/xor/api/pkg/game"
)

func ResetQuiz(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		appState := ctx.GetContextVar[*config.AppState](r.Context(), ctx.StateKey)
		game.SelectAQuiz(appState)
		next.ServeHTTP(w, r)
	})
}
