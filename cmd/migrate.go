package cmd

import (
	"github.com/amirhnajafiz/authX/internal/config"
	"github.com/amirhnajafiz/authX/internal/model"
	"github.com/amirhnajafiz/authX/internal/storage"

	"go.uber.org/zap"
)

// Migrate command.
type Migrate struct {
	Cfg    config.Config
	Logger *zap.Logger
}

// main function of Migrate command.
func (m Migrate) main() {
	// open db
	db, err := storage.NewConnection(m.Cfg.Storage)
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
