package database

import (
	"fmt"

	"github.com/yuttasakcom/go-hexa/app/config"
	"github.com/yuttasakcom/go-hexa/app/todo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormStore struct {
	db *gorm.DB
}

func NewGormStore(db *gorm.DB) *GormStore {
	return &GormStore{db: db}
}

func (s *GormStore) Create(todo *todo.Todo) error {
	return s.db.Create(todo).Error
}

func GormStoreConnect(config config.IDb) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.Url()), &gorm.Config{})
	if err != nil {
		fmt.Println("db err: (Init) ", err)
	}
	return db
}
