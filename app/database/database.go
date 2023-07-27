package database

import (
	"github.com/yuttasakcom/go-hexa/app/config"
	"go.mongodb.org/mongo-driver/mongo"
)

func DatabaseConnect(config config.IDb) *Store {
	// return NewStore(GormConnect(config))
	return NewStore(MongoConnect(config))
}

type Store struct {
	// *GormDB
	*MongoDB
}

// func NewStore(db *gorm.DB) *Store {
// 	return &Store{&GormDB{db}}
// }

func NewStore(db *mongo.Database) *Store {
	return &Store{&MongoDB{db}}
}
