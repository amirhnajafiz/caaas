package apps

import (
	"github.com/amirhnajafiz/authX/internal/model"

	"gorm.io/gorm"
)

type Apps interface {
	Create(app *model.App) error
	GetByKey(key string) (*model.App, error)
	GetByUserID(userID uint) ([]*model.App, error)
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

func (a *apps) GetByKey(key string) (*model.App, error) {
	app := new(model.App)

	if err := a.db.First(&app).Where("app_key = ?", key).Error; err != nil {
		return nil, err
	}

	return app, nil
}

func (a *apps) GetByUserID(userID uint) ([]*model.App, error) {
	var list []*model.App

	if err := a.db.Find(&list).Where("user_id = ?", userID).Error; err != nil {
		return nil, err
	}

	return list, nil
}
