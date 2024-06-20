package handlers

import (
	"github.com/amirhnajafiz/caaas/internal/config"
	"github.com/amirhnajafiz/caaas/internal/controller"
	"github.com/amirhnajafiz/caaas/internal/handlers/migrate"
	"github.com/amirhnajafiz/caaas/pkg/enum"
	"github.com/amirhnajafiz/caaas/pkg/jwt"

	"github.com/go-pg/pg/v10"
	"go.uber.org/zap"
)

// loader is a struct that holds import components of our handlers.
type loader struct {
	cfg      config.Config
	auth     *jwt.Auth
	logger   *zap.Logger
	ctl      *controller.Controller
	database *pg.DB
}

// bootstrap method is use for initializing base components.
func (l *loader) bootstrap() {

}

func LoadHandler(cfg config.Config, db *pg.DB) Handler {
	l := loader{
		cfg:      cfg,
		database: db,
	}
	l.bootstrap()

	switch cfg.Mode {
	case enum.ModeAPI:
		return nil
	case enum.ModeGW:
		return nil
	case enum.ModeMigrate:
		return &migrate.Handler{
			Database: l.database,
		}
	default:
		return nil
	}
}
