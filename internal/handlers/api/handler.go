package api

import (
	"fmt"

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

func (h Handler) Execute() error {
	e := echo.New()

	return e.Start(fmt.Sprintf(":%d", h.Port))
}
