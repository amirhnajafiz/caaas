package main

import (
	"log"

	"github.com/amirhnajafiz/caaas/internal/config"
	"github.com/amirhnajafiz/caaas/internal/handlers"
	"github.com/amirhnajafiz/caaas/internal/storage"
)

func main() {
	// load configs
	cfg := config.LoadConfigs()

	// open database connection
	db, err := storage.NewConnection(cfg.Storage)
	if err != nil {
		log.Fatal(err)
	}

	// reload a handler and execute it
	if err := handlers.LoadHandler(cfg, db).Execute(); err != nil {
		log.Fatal(err)
	}
}
