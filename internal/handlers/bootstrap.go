package handlers

import (
	"github.com/amirhnajafiz/caaas/internal/config"
	"github.com/amirhnajafiz/caaas/internal/controller"
	"github.com/amirhnajafiz/caaas/internal/enum"
	"github.com/amirhnajafiz/caaas/internal/handlers/api"
	"github.com/amirhnajafiz/caaas/internal/handlers/gateway"
	"github.com/amirhnajafiz/caaas/internal/handlers/migrate"
	"github.com/amirhnajafiz/caaas/internal/monitoring/logger"
	"github.com/amirhnajafiz/caaas/internal/monitoring/metrics"
	"github.com/amirhnajafiz/caaas/pkg/jwt"

	"github.com/go-pg/pg/v10"
	"go.uber.org/zap"
)

// loader is a struct that holds import components of our handlers.
type loader struct {
	cfg      config.Config
	metrics  *metrics.Metrics
	auth     *jwt.Auth
	logger   *zap.Logger
	ctl      *controller.Controller
	database *pg.DB
}

// bootstrap method is use for initializing base components.
func (l *loader) bootstrap() {
	// create a new controller
	l.ctl = controller.NewController(l.database)
	// create a new logger
	l.logger = logger.NewLogger(l.cfg.Logger)
	// create a new auth system
	l.auth = jwt.New(l.cfg.Auth)
	// create metrics instance
	l.metrics = metrics.NewMetrics()
	// create metrics server if enable
	if l.cfg.Mode == enum.ModeGW && l.cfg.Metrics.Enable {
		metrics.StartServer(l.cfg.Metrics.Port)
	}
}

// LoadHandler returns a handler based on the mode which is set in configs.
func LoadHandler(cfg config.Config, db *pg.DB) Handler {
	// bootstrap section
	l := loader{
		cfg:      cfg,
		database: db,
	}
	l.bootstrap()

	// handler selector
	switch cfg.Mode {
	case enum.ModeAPI:
		return &api.Handler{
			Logger: l.logger.Named("api"),
			Ctl:    l.ctl,
			Port:   l.cfg.HTTPServerPort,
		}
	case enum.ModeGW:
		return &gateway.Handler{
			Logger:  l.logger.Named("gateway"),
			Ctl:     l.ctl,
			Auth:    l.auth,
			Port:    l.cfg.HTTPServerPort,
			Metrics: l.metrics,
		}
	case enum.ModeMigrate:
		return &migrate.Handler{
			Database: l.database,
		}
	default:
		return nil
	}
}
