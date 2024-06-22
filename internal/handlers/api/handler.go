package api

import (
	"fmt"
	"net/http"

	"github.com/amirhnajafiz/caaas/internal/controller"

	"github.com/labstack/echo/v4"
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
	v1 := api.Group("/v1")
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

	return e.Start(fmt.Sprintf(":%d", h.Port))
}
