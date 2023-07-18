package main

import (
	"log"
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
	// Liveness Probe
	_, err := os.Create("/tmp/live")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove("/tmp/live")

	cfg := config.NewConfig(envFile())
	store := database.DatabaseConnect(cfg.PgDB())
	server.NewServer(cfg, store).Start()
}
