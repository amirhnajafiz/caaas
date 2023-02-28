package middleware

import (
	"github.com/amirhnajafiz/authX/internal/repository"
)

type Middleware struct {
	Repository repository.Repository
}
