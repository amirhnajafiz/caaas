package apps

import (
	"github.com/amirhnajafiz/authX/internal/model"

	"gorm.io/gorm"
)

type Apps interface {
	Create(app *model.App) error
	GetSingle(id uint) (*model.App, error)
	Delete(id uint) error
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
	return nil
}

func (a *apps) GetSingle(id uint) (*model.App, error) {
	return nil, nil
}

func (a *apps) Delete(id uint) error {
	return nil
}
