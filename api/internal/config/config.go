package config

import (
	"github.com/dyrector-io/xor/api/pkg/processor"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type AppConfig struct {
	AllowedOrigins string `env:"ALLOWED_ORIGINS"`
	Port           uint16 `env:"PORT" env-default:"3333"`
	Debug          bool   `env:"DEBUG"`
	DSN            string `env:"DSN"`
	Method         string `env:"METHOD"`
	Freq           string `env:"FREQ" env-default:"DAILY"`
	EndDate        string `env:"END_DATE"`
}
type AppState struct {
	AppConfig   *AppConfig
	DBConn      *gorm.DB
	Ended       bool
	QuizList    processor.QuizSequence
	QuizCounter int
}

func (c *AppConfig) MarshalZerologObject(e *zerolog.Event) {
	e.Str("origins", c.AllowedOrigins).
		Uint16("port", c.Port).
		Bool("debug", c.Debug).
		Str("method", c.Method).
		Str("freq", c.Freq).
		Str("end", c.EndDate)
}
