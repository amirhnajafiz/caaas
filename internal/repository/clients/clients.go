package clients

import (
	"github.com/amirhnajafiz/authX/internal/model"
	"gorm.io/gorm"
)

type Clients interface {
	Create(client *model.Client) error
	GetAppClients(key string) ([]uint, error)
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

func (a *clients) GetAppClients(key string) ([]uint, error) {
	var (
		list []*model.Client
		ids  []uint
	)

	if err := a.db.Find(&list).Where("app_key = ?", key).Error; err != nil {
		return nil, err
	}

	for _, item := range list {
		ids = append(ids, item.ID)
	}

	return ids, nil
}

func (a *clients) GetSingle(id uint) (*model.Client, error) {
	client := new(model.Client)

	if err := a.db.Table("clients").Find(&client).Where("id = ?", id).Error; err != nil {
		return nil, err
	}

	return client, nil
}
