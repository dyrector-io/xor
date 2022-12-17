package db

import (
	"gorm.io/gorm"
)

type Pick struct {
	gorm.Model
	Picks []int
}

func PersistPicks(picks []int) error {
	return nil
}

func InitPostgres() {}
