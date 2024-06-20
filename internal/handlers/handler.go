package handlers

import (
	"github.com/amirhnajafiz/caaas/internal/controller"
	"github.com/amirhnajafiz/caaas/internal/handlers/migrate"
	"github.com/amirhnajafiz/caaas/pkg/enum"
	"github.com/go-pg/pg/v10"
)

// Handler is a abstract struct of api, gateway, and migrate handers,
// which will be used based on the give mode.
type Handler interface {
	Execute() error
}

func LoadHandler(mode string, db *pg.DB) Handler {
	_ = controller.NewController(db)

	switch mode {
	case enum.ModeAPI:
		return nil
	case enum.ModeGW:
		return nil
	case enum.ModeMigrate:
		return &migrate.Handler{
			Database: db,
		}
	default:
		return nil
	}
}
