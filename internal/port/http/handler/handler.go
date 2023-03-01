package handler

import (
	"github.com/amirhnajafiz/authX/internal/repository"

	"go.uber.org/zap"
)

type Handler struct {
	Logger     *zap.Logger
	Repository repository.Repository
}
