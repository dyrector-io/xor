package config

import (
	"github.com/dyrector-io/xor/api/pkg/processor"
	"gorm.io/gorm"
)

type AppConfig struct {
	AllowedOrigins string `env:"ALLOWED_ORIGINS"`
	Port           uint16 `env:"PORT" env-default:"3333"`
	Debug          bool   `env:"DEBUG"`
	DSN            string `env:"DSN"`
	Method         string `env:"METHOD"`
}
type AppState struct {
	AppConfig *AppConfig
	DBConn    *gorm.DB
	QuizList  processor.CNCFSequence
}
