package v2

import (
	"github.com/amirhnajafiz/caaas/internal/controller"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// API handler is a private handler to be used for
// managing users and groups.
type Handler struct {
	Logger *zap.Logger
	Ctl    *controller.Controller
}

func (h Handler) New(v2 *echo.Group) {
	users := v2.Group("/users")
	roles := v2.Group("/roles")

	// users methods
	users.GET("/", h.getUser)

	// roles methods
	roles.POST("/", h.addRoleToUser)
	roles.PATCH("/", h.removeRoleFromUser)
	roles.DELETE("/", h.removeRole)
}
