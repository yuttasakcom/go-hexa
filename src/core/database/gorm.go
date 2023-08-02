package database

import (
	slog "github.com/Sellsuki/sellsuki-go-logger"
	"github.com/yuttasakcom/go-kafka-simple/src/core/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormDB struct {
	*gorm.DB
}

func GormConnect(pgCfg config.DBConfiger) *gorm.DB {
	db, err := gorm.Open(postgres.Open(pgCfg.Url()), &gorm.Config{})
	if err != nil {
		slog.L().Fatal("Error connect postgres", slog.Error(err))
	}
	slog.L().Info("Connected to Postgres!")
	return db
}

func (g *GormDB) Create(v interface{}) error {
	return g.DB.Create(v).Error
}
