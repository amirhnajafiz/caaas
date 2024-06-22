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

	// register k8s endpoints
	e.GET("/healthz", h.health)

	// register metric needed enpoints
	counts := e.Group("", h.requestsMiddleware)

	// loging endpoint
	counts.POST("/", h.login)

	// user endpoints
	auth := counts.Group("", h.authMiddleware)
	auth.GET("/", h.validate)
	auth.GET("/groups", h.groups)

	return e.Start(fmt.Sprintf(":%d", h.Port))
}
