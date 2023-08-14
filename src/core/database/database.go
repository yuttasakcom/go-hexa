package database

import (
	"github.com/yuttasakcom/go-hexa/src/core/config"
)

type Store struct {
	// Use MongoDB
	*MongoDB

	// Use PostgreSQL
	// *GormDB
}

// Easy to change database
func NewStore(config config.DBConfig) *Store {
	// Use MongoDB
	return &Store{&MongoDB{(MongoConnect(config.Mg))}}

	// Use PostgreSQL
	// return &Store{&GormDB{(GormConnect(config.Pg))}}
}
