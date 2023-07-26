package config

import (
	"fmt"
	"log"
	"strconv"

	"github.com/joho/godotenv"
)

type config struct {
	app  *app
	pgDb *pgDB
}

type IConfig interface {
	App() *app
	PgDB() *pgDB
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
		pgDb: &pgDB{
			host: env["PG_DB_HOST"],
			port: func() int {
				p, err := strconv.Atoi(env["PG_DB_PORT"])
				if err != nil {
					log.Fatalf("Error port fail %v", err)
				}
				return p
			}(),
			user:     env["PG_DB_USER"],
			password: env["PG_DB_PASSWORD"],
			dbname:   env["PG_DB_NAME"],
			sslmode:  env["PG_DB_SSLMODE"],
			timezone: env["PG_DB_TIMEZONE"],
		},
	}
}

func (c *config) App() *app {
	return c.app
}

func (c *config) PgDB() *pgDB {
	return c.pgDb
}

type app struct {
	host string
	port int
}

func (a *app) Host() string {
	return fmt.Sprintf("%s:%d", a.host, a.port)
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

type IDb interface {
	Url() string
}

func (d *pgDB) Url() string {
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
