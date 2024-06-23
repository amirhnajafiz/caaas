package v1

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

func (h Handler) New(v1 *echo.Group) {
	users := v1.Group("/users")
	groups := v1.Group("/groups")

	// users methods
	users.GET("/", h.getAllUsers)
	users.POST("/", h.createUser)
	users.PATCH("/", h.updateUser)
	users.DELETE("/", h.removeUser)

	// groups methods
	groups.POST("/", h.addUserToGroup)
	groups.PATCH("/", h.removeUserFromGroup)
	groups.DELETE("/", h.removeGroup)
}
