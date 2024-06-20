package api

import (
	"github.com/amirhnajafiz/caaas/internal/controller"

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
	return nil
}
