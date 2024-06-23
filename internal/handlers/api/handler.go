package api

import (
	"fmt"
	"net/http"

	"github.com/amirhnajafiz/caaas/internal/controller"
	v1 "github.com/amirhnajafiz/caaas/internal/handlers/api/v1"
	v2 "github.com/amirhnajafiz/caaas/internal/handlers/api/v2"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

// API handler is a private handler to be used for
// managing users and groups.
type Handler struct {
	Logger *zap.Logger
	Ctl    *controller.Controller
	Port   int
}

func (h Handler) health(c echo.Context) error {
	return c.String(http.StatusOK, "")
}

func (h Handler) Execute() error {
	e := echo.New()

	// register k8s endpoints
	e.GET("/healthz", h.health)

	// register normal endpoints
	api := e.Group("/api")

	// using logger middleware for api
	api.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			h.Logger.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)

			return nil
		},
	}))

	// create all v.hanlders
	v1.Handler{
		Logger: h.Logger.Named("v1"),
		Ctl:    h.Ctl,
	}.New(api.Group("/v1"))
	v2.Handler{
		Logger: h.Logger.Named("v2"),
		Ctl:    h.Ctl,
	}.New(api.Group("/v2"))

	return e.Start(fmt.Sprintf(":%d", h.Port))
}
