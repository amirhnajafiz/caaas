package middleware

import (
	"github.com/amirhnajafiz/authX/internal/repository"
	"github.com/amirhnajafiz/authX/pkg/auth"

	"go.uber.org/zap"
)

type Middleware struct {
	Auth       auth.Auth
	Logger     *zap.Logger
	Repository repository.Repository
}
