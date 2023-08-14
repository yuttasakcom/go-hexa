package config

import (
	"os"
	"strconv"

	slog "github.com/Sellsuki/sellsuki-go-logger"
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
		slog.L().Warn("Error loading .env file: %v", err)
	}

	return config{
		app: App{
			host: os.Getenv("APP_HOST"),
			port: func() int {
				p, err := strconv.Atoi(os.Getenv("APP_PORT"))
				if err != nil {
					slog.L().Warn("Error converting APP_PORT to int: %v", err)
				}
				return p
			}(),
			DebugLog: func() bool {
				d, err := strconv.ParseBool(os.Getenv("APP_DEBUG_LOG"))
				if err != nil {
					slog.L().Warn("Error converting APP_DEBUG_LOG to bool: %v", err)
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
						slog.L().Warn("Error PG_DB_PORT fail %v", err)
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
						slog.L().Warn("Error MG_DB_PORT fail %v", err)
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
					slog.L().Warn("Error JAEGER_PORT fail %v", err)
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
