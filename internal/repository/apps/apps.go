package apps

import (
	"github.com/amirhnajafiz/authX/internal/model"

	"gorm.io/gorm"
)

// Apps manages the apps model.
type Apps interface {
	Create(app *model.App) error
}

// New generates new app repository.
func New(db *gorm.DB) Apps {
	return &apps{
		db: db,
	}
}

// apps manages the functions of repository.
type apps struct {
	db *gorm.DB
}

// Create a new app.
func (a *apps) Create(app *model.App) error {
	return a.db.Create(app).Error
}
