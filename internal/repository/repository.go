package repository

import (
	"github.com/amirhnajafiz/authX/internal/repository/apps"
	"github.com/amirhnajafiz/authX/internal/repository/clients"
	"github.com/amirhnajafiz/authX/internal/repository/users"

	"gorm.io/gorm"
)

// Repository manages the database entities.
type Repository struct {
	Apps    apps.Apps
	Clients clients.Clients
	Users   users.Users
}

// New repository.
func New(db *gorm.DB) Repository {
	return Repository{
		Apps:    apps.New(db),
		Clients: clients.New(db),
		Users:   users.New(db),
	}
}
