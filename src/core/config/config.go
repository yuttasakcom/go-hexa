package config

import (
	"log"
	"os"
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
	err := godotenv.Load(path)
	if err != nil {
		log.Printf("Error loading %v\n", path)
	}

	return config{
		app: App{
			host: os.Getenv("APP_HOST"),
			port: func() int {
				p, err := strconv.Atoi(os.Getenv("APP_PORT"))
				if err != nil {
					log.Printf("Error converting APP_PORT to int: %v\n", err)
				}
				return p
			}(),
			grpc_port: func() int {
				p, err := strconv.Atoi(os.Getenv("APP_GRPC_PORT"))
				if err != nil {
					log.Printf("Error converting APP_GRPC_PORT to int: %v\n", err)
				}
				return p
			}(),
			DebugLog: func() bool {
				d, err := strconv.ParseBool(os.Getenv("APP_DEBUG_LOG"))
				if err != nil {
					log.Printf("Error converting APP_DEBUG_LOG to bool: %v\n", err)
				}
				return d
			}(),
			AppName:    os.Getenv("APP_NAME"),
			AppVersion: os.Getenv("APP_VERSION"),
			AppEnv:     os.Getenv("APP_ENV"),
		},
		dbConfig: DBConfig{
			Pg: pgDB{
				host: os.Getenv("PG_DB_HOST"),
				port: func() int {
					p, err := strconv.Atoi(os.Getenv("PG_DB_PORT"))
					if err != nil {
						log.Printf("Error PG_DB_PORT fail %v\n", err)
					}
					return p
				}(),
				user:     os.Getenv("PG_DB_USER"),
				password: os.Getenv("PG_DB_PASSWORD"),
				dbname:   os.Getenv("PG_DB_NAME"),
				sslmode:  os.Getenv("PG_DB_SSLMODE"),
				timezone: os.Getenv("PG_DB_TIMEZONE"),
			},
			Mg: mgDB{
				host: os.Getenv("MG_DB_HOST"),
				port: func() int {
					p, err := strconv.Atoi(os.Getenv("MG_DB_PORT"))
					if err != nil {
						log.Printf("Error MG_DB_PORT fail %v\n", err)
					}
					return p
				}(),
				user:     os.Getenv("MG_DB_USER"),
				password: os.Getenv("MG_DB_PASSWORD"),
				dbname:   os.Getenv("MG_DB_NAME"),
			},
		},
		jaeger: Jaeger{
			host: os.Getenv("JAEGER_HOST"),
			port: func() int {
				p, err := strconv.Atoi(os.Getenv("JAEGER_PORT"))
				if err != nil {
					log.Printf("Error JAEGER_PORT fail %v\n", err)
				}
				return p
			}(),
			uri: os.Getenv("JAEGER_URI"),
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
