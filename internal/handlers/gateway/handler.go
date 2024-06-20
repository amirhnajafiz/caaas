package gateway

import (
	"github.com/amirhnajafiz/caaas/internal/controller"
	"github.com/amirhnajafiz/caaas/pkg/jwt"

	"go.uber.org/zap"
)

// Gateway handler is responsible for handling users
// authentication and authorization requests.
type Handler struct {
	Logger *zap.Logger
	Ctl    *controller.Controller
	Auth   *jwt.Auth
	Port   int
}

func (h Handler) Execute() error {
	return nil
}
