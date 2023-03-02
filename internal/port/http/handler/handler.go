package handler

import (
	"github.com/amirhnajafiz/authX/internal/repository"
	"github.com/amirhnajafiz/authX/pkg/auth"

	"go.uber.org/zap"
)

type Handler struct {
	Auth       *auth.Auth
	Logger     *zap.Logger
	Repository repository.Repository
}
