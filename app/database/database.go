package database

import (
	"github.com/yuttasakcom/go-hexa/app/config"
	"go.mongodb.org/mongo-driver/mongo"
)

func DatabaseConnect(config config.IDb) *Store {
	// return NewGormStore(GormConnect(config))
	return NewMongoStore(MongoConnect(config))
}

type Store struct {
	// *gorm.DB
	*mongo.Database
}

// func NewGormStore(db *gorm.DB) *Store {
// 	return &Store{DB: db}
// }

func NewMongoStore(db *mongo.Database) *Store {
	return &Store{Database: db}
}
