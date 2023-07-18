package main

import (
	"os"

	"github.com/yuttasakcom/go-hexa/app/config"
	"github.com/yuttasakcom/go-hexa/app/database"
	"github.com/yuttasakcom/go-hexa/app/server"
)

func envFile() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}

func main() {
	cfg := config.NewConfig(envFile())
	store := database.DatabaseConnect(cfg.PgDB())
	server.NewServer(cfg, store).Start()
}
