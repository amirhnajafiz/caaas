package apps

import (
	"github.com/amirhnajafiz/authX/internal/model"

	"gorm.io/gorm"
)

type Apps interface {
	Create(app *model.App) error
	GetSingle(id uint) (*model.App, error)
}

func New(db *gorm.DB) Apps {
	return &apps{
		db: db,
	}
}

type apps struct {
	db *gorm.DB
}

func (a *apps) Create(app *model.App) error {
	return a.db.Create(app).Error
}

func (a *apps) GetSingle(id uint) (*model.App, error) {
	app := new(model.App)

	if err := a.db.First(&app).Where("id = ?", id).Error; err != nil {
		return nil, err
	}

	return app, nil
}
