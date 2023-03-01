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
		m.Logger.Error("database connection failed", zap.Error(err))

		return
	}

	// migrate into database
	if er := db.AutoMigrate(
		&model.App{},
		&model.User{},
		&model.Client{},
	); er != nil {
		m.Logger.Error("migration failed", zap.Error(er))
	}
}
