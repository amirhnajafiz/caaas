package clients

import (
	"github.com/amirhnajafiz/authX/internal/model"

	"gorm.io/gorm"
)

type Clients interface {
	Create(client *model.Client) error
	GetAppClients(id uint) ([]uint, error)
	GetSingle(id uint) (*model.Client, error)
}

func New(db *gorm.DB) Clients {
	return &clients{
		db: db,
	}
}

type clients struct {
	db *gorm.DB
}

func (a *clients) Create(client *model.Client) error {
	return a.db.Create(client).Error
}

func (a *clients) GetAppClients(id uint) ([]uint, error) {
	var list []uint

	if err := a.db.Raw("select id from clients where app_id = ?", id).Find(list).Error; err != nil {
		return nil, err
	}

	return list, nil
}

func (a *clients) GetSingle(id uint) (*model.Client, error) {
	client := new(model.Client)

	if err := a.db.Find(client).Where("id = ?", id).Error; err != nil {
		return nil, err
	}

	return client, nil
}
