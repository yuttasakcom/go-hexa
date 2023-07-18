package database

import (
	"fmt"

	"github.com/yuttasakcom/go-hexa/app/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GormConnect(config config.IDb) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.Url()), &gorm.Config{})
	if err != nil {
		fmt.Println("db err: (Init) ", err)
	}
	return db
}
