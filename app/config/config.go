package config

import (
	"fmt"
	"log"
	"strconv"

	"github.com/joho/godotenv"
)

type config struct {
	app *app
	db  *db
}

type IConfig interface {
	App() *app
	Db() *db
}

func NewConfig(path string) IConfig {
	env, err := godotenv.Read(path)
	if err != nil {
		log.Fatalf("Error dotenv fail %v", err)
	}
	return &config{
		app: &app{
			host: env["APP_HOST"],
			port: func() int {
				p, err := strconv.Atoi(env["APP_PORT"])
				if err != nil {
					log.Fatalf("Error port fail %v", err)
				}
				return p
			}(),
		},
		db: &db{
			host: env["DB_HOST"],
			port: func() int {
				p, err := strconv.Atoi(env["DB_PORT"])
				if err != nil {
					log.Fatalf("Error port fail %v", err)
				}
				return p
			}(),
			user:     env["DB_USER"],
			password: env["DB_PASSWORD"],
			dbname:   env["DB_NAME"],
			sslmode:  env["DB_SSLMODE"],
			timezone: env["DB_TIMEZONE"],
		},
	}
}

func (c *config) App() *app {
	return c.app
}

func (c *config) Db() *db {
	return c.db
}

type app struct {
	host string
	port int
}

func (a *app) Host() string {
	return fmt.Sprintf("%s:%d", a.host, a.port)
}

type db struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
	sslmode  string
	timezone string
}

type IDb interface {
	Url() string
}

func (d *db) Url() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		d.host,
		d.port,
		d.user,
		d.password,
		d.dbname,
		d.sslmode,
	)
}
