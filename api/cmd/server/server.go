package main

import (
	"os"

	"github.com/dyrector-io/xor/api/internal/config"
	"github.com/dyrector-io/xor/api/pkg/server"
	"github.com/rs/zerolog/log"

	"github.com/ilyakaznacheev/cleanenv"
)

func ReadConfig(cfg *config.AppConfig) error {
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

func main() {
	appConfig := &config.AppConfig{}

	err := ReadConfig(appConfig)
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	serv := server.GetChi(appConfig)

	log.Info().Msgf("starting server at: %d", appConfig.Port)
	err = serv.ListenAndServe()
	if err != nil {
		log.Fatal().Err(err).Send()
	}
}
