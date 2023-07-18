package database

import (
	"github.com/yuttasakcom/go-hexa/app/config"
	"gorm.io/gorm"
)

func DatabaseConnect(config config.IDb) *Store {
	return NewStore(GormConnect(config))
}

type Store struct {
	*gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{DB: db}
}
