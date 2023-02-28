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
	return nil
}

func (a *clients) GetAppClients(id uint) ([]uint, error) {
	return nil, nil
}

func (a *clients) GetSingle(id uint) (*model.Client, error) {
	return nil, nil
}
