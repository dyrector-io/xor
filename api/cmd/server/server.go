package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/dyrector-io/xor/api/pkg/game"
	"github.com/rs/zerolog/log"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/ilyakaznacheev/cleanenv"
)

const (
	HTTPReadHeaderTimeout = 3 * time.Second
	HTTPTimeout           = 60 * time.Second
)

type AppConfig struct {
	PORT uint16 `env:"PORT" env-default:"3333"`
}

type AppState struct {
	// DBConnection db ``
}

func ReadConfig(cfg *AppConfig) error {
	err := cleanenv.ReadConfig(".env", cfg)

	if err != nil && !os.IsNotExist(err) {
		return err
	}
	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return err
	}

	return nil
}

type LogWriter struct {
	http.ResponseWriter
}

func (w LogWriter) Write(p []byte) {
	_, err := w.ResponseWriter.Write(p)
	if err != nil {
		log.Error().Err(err).Msgf("Write failed: %v", err)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(HTTPTimeout))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		LogWriter{w}.Write([]byte("ok"))
	})

	r.Get("/quiz", game.GetQuiz)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		LogWriter{w}.Write([]byte("pong"))
	})

	r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("test")
	})

	appConfig := &AppConfig{}

	err := ReadConfig(appConfig)
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	server := &http.Server{
		Addr:              fmt.Sprintf(":%v", appConfig.PORT),
		ReadHeaderTimeout: HTTPReadHeaderTimeout,
		Handler:           r,
	}

	log.Info().Msgf("starting server at: %d", appConfig.PORT)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal().Err(err).Send()
	}
}
