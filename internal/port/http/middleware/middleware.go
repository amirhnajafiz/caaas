package middleware

import (
	"github.com/amirhnajafiz/authX/internal/repository"
	"github.com/amirhnajafiz/authX/pkg/auth"
)

type Middleware struct {
	Auth       auth.Auth
	Repository repository.Repository
}
