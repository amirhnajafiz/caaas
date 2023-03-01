package cmd

import (
	"github.com/amirhnajafiz/authX/internal/config"
	"github.com/amirhnajafiz/authX/internal/model"
	"github.com/amirhnajafiz/authX/internal/storage"
)

// Migrate command.
type Migrate struct{}

// main function of Migrate command.
func (m Migrate) main() {
	// load configs
	cfg := config.LoadConfigs()

	// open db
	db, err := storage.NewConnection(cfg.Storage)
	if err != nil {
		panic(err)
	}

	// migrate into database
	if err := db.AutoMigrate(
		&model.App{},
		&model.User{},
		&model.Client{},
	); err != nil {
		panic(err)
	}
}
