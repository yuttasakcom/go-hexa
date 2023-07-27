package database

import (
	"fmt"

	"github.com/yuttasakcom/go-hexa/app/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormDB struct {
	*gorm.DB
}

func GormConnect(config config.IDb) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.Url()), &gorm.Config{})
	if err != nil {
		fmt.Println("db err: (Init) ", err)
	}
	fmt.Println("Connected to Postgres!")
	return db
}

func (g *GormDB) Create(v interface{}) error {
	return g.DB.Create(v).Error
}
