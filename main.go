package main

import (
	"log"

	"github.com/amirhnajafiz/caaas/internal/config"
	"github.com/amirhnajafiz/caaas/internal/handlers"
	"github.com/amirhnajafiz/caaas/internal/storage"
)

func main() {
	// load configs
	cfg := config.LoadConfigs("config.yaml")

	// open database connection
	db, err := storage.NewConnection(cfg.Storage)
	if err != nil {
		log.Fatal(err)
	}

	// reload a handler and execute it
	hd := handlers.LoadHandler(cfg, db)
	if hd == nil {
		log.Fatalf("give mode is not supported\n")
	}
	if err := hd.Execute(); err != nil {
		log.Fatal(err)
	}
}
