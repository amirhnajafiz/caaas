package apps

import (
	"fmt"

	"github.com/amirhnajafiz/authX/internal/model"

	"gorm.io/gorm"
)

// Apps manages the apps model.
type Apps interface {
	Create(app *model.App) error
	Get(userID uint) (*model.App, error)
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

// Get an app by user id.
func (a *apps) Get(userID uint) (*model.App, error) {
	app := new(model.App)

	if err := a.db.Where("user_id = ?", userID).Find(&app).Error; err != nil {
		return nil, err
	}

	// app not found
	if app.ID == 0 {
		return nil, fmt.Errorf("app not found")
	}

	return app, nil
}
