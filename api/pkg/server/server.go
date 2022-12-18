package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dyrector-io/xor/api/internal/config"
	"github.com/dyrector-io/xor/api/internal/ctx"
	"github.com/dyrector-io/xor/api/pkg/database"
	"github.com/dyrector-io/xor/api/pkg/game"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
)

const (
	HTTPReadHeaderTimeout = 3 * time.Second
	HTTPTimeout           = 60 * time.Second
)

type LogWriter struct {
	http.ResponseWriter
}

func (w LogWriter) Write(p []byte) {
	_, err := w.ResponseWriter.Write(p)
	if err != nil {
		log.Error().Err(err).Msgf("write failed: %v", err)
	}
}

func GetChi(appConfig *config.AppConfig) *http.Server {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	db := database.InitPostgres(appConfig)
	appState := &config.AppState{
		AppConfig: appConfig,
		DBConn:    db,
	}

	today := time.Now()
	game.SelectAQuiz(appState, today)

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r.WithContext(ctx.SetContextVar(r.Context(), ctx.StateKey, appState)))
		})
	})

	r.Use(middleware.Timeout(HTTPTimeout))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		LogWriter{w}.WriteHeader(http.StatusOK)
	})

	r.Get("/quiz", GetQuiz)

	r.Get("/history", GetHistory)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		LogWriter{w}.Write([]byte("pong"))
	})

	r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("test")
	})

	server := &http.Server{
		Addr:              fmt.Sprintf(":%v", appConfig.Port),
		ReadHeaderTimeout: HTTPReadHeaderTimeout,
		Handler:           r,
	}

	s := gocron.NewScheduler(time.UTC)
	_, err := s.Every(1).Minute().Do(game.SelectAQuiz, appState, today)
	if err != nil {
		log.Error().Err(err).Msg("daily quiz gen cron is not ok")
	}
	s.StartAsync()

	return server
}
