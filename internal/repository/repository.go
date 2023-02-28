package repository

import (
	"github.com/amirhnajafiz/authX/internal/repository/apps"
	"github.com/amirhnajafiz/authX/internal/repository/clients"
	"github.com/amirhnajafiz/authX/internal/repository/users"
)

// Repository manages the database entities.
type Repository struct {
	Apps    apps.Apps
	Clients clients.Clients
	Users   users.Users
}

// New repository.
func New() Repository {
	return Repository{
		Apps:    apps.New(),
		Clients: clients.New(),
		Users:   users.New(),
	}
}
