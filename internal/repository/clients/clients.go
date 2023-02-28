package clients

import (
	"github.com/amirhnajafiz/authX/internal/model"
)

type Clients interface {
	Create(client *model.Client) error
	GetAppClients(id uint) ([]uint, error)
	GetSingle(id uint) (*model.Client, error)
}

func New() Clients {
	return &clients{}
}

type clients struct{}

func (a *clients) Create(client *model.Client) error {
	return nil
}

func (a *clients) GetAppClients(id uint) ([]uint, error) {
	return nil, nil
}

func (a *clients) GetSingle(id uint) (*model.Client, error) {
	return nil, nil
}
