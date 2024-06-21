package gateway

import (
	"fmt"
	"net/http"

	"github.com/amirhnajafiz/caaas/internal/controller"
	"github.com/amirhnajafiz/caaas/internal/monitoring/metrics"
	"github.com/amirhnajafiz/caaas/pkg/jwt"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// Gateway handler is responsible for handling users
// authentication and authorization requests.
type Handler struct {
	Logger  *zap.Logger
	Ctl     *controller.Controller
	Auth    *jwt.Auth
	Metrics *metrics.Metrics
	Port    int
}

func (h Handler) health(c echo.Context) error {
	return c.String(http.StatusOK, "")
}

func (h Handler) Execute() error {
	e := echo.New()

	// register endpoints
	e.GET("/healthz", h.health)

	// register metric needed enpoints
	e.Use(h.requestsMiddleware)

	return e.Start(fmt.Sprintf(":%d", h.Port))
}
