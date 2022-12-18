package config

import (
	"github.com/dyrector-io/xor/api/pkg/processor"
	"gorm.io/gorm"
)

type AppConfig struct {
	Port   uint16 `env:"PORT" env-default:"3333"`
	Debug  bool   `env:"DEBUG"`
	DSN    string `env:"DSN"`
	Method string `env:"METHOD" env-default:"RANDOM"`
}
type AppState struct {
	AppConfig *AppConfig
	DBConn    *gorm.DB
	QuizList  processor.CNCFSequence
}
