package config

import (
	"log"
	"strconv"

	"github.com/joho/godotenv"
)

type Configer interface {
	App() App
	DBConfig() DBConfig
	Jaeger() Jaeger
}
type config struct {
	app      App
	dbConfig DBConfig
	jaeger   Jaeger
}

func NewConfig(path string) config {
	env, err := godotenv.Read(path)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return config{
		app: App{
			host: env["APP_HOST"],
			port: func() int {
				p, err := strconv.Atoi(env["APP_PORT"])
				if err != nil {
					log.Fatalf("Error converting APP_PORT to int: %v", err)
				}
				return p
			}(),
			DebugLog: func() bool {
				d, err := strconv.ParseBool(env["APP_DEBUG_LOG"])
				if err != nil {
					log.Fatalf("Error converting APP_DEBUG_LOG to bool: %v", err)
				}
				return d
			}(),
			AppName:    env["APP_NAME"],
			AppVersion: env["APP_VERSION"],
			AppEnv:     env["APP_ENV"],
		},
		dbConfig: DBConfig{
			Pg: pgDB{
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
			Mg: mgDB{
				host: env["MG_DB_HOST"],
				port: func() int {
					p, err := strconv.Atoi(env["MG_DB_PORT"])
					if err != nil {
						log.Fatalf("Error port fail %v", err)
					}
					return p
				}(),
				user:     env["MG_DB_USER"],
				password: env["MG_DB_PASSWORD"],
				dbname:   env["MG_DB_NAME"],
			},
		},
		jaeger: Jaeger{
			host: env["JAEGER_HOST"],
			port: func() int {
				p, err := strconv.Atoi(env["JAEGER_PORT"])
				if err != nil {
					log.Fatalf("Error port fail %v", err)
				}
				return p
			}(),
			uri: env["JAEGER_URI"],
		},
	}
}

func (c config) App() App {
	return c.app
}

func (c config) DBConfig() DBConfig {
	return c.dbConfig
}

func (c config) Jaeger() Jaeger {
	return c.jaeger
}
