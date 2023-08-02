package config

import "fmt"

// DBConfig is the struct that contains the database configuration
type DBConfig struct {
	Pg pgDB
	Mg mgDB
}

type DBConfiger interface {
	Url() string
	Dbname() string
}

type pgDB struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
	sslmode  string
	timezone string
}

type mgDB struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
}

func (d pgDB) Url() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		d.host,
		d.port,
		d.user,
		d.password,
		d.dbname,
		d.sslmode,
		d.timezone,
	)
}

func (d pgDB) Dbname() string {
	return d.dbname
}

func (d mgDB) Url() string {
	return fmt.Sprintf(
		"mongodb://%s:%s@%s:%d/%s?ssl=false&authSource=admin",
		d.user,
		d.password,
		d.host,
		d.port,
		d.dbname,
	)
}

func (d mgDB) Dbname() string {
	return d.dbname
}
