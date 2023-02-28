package users

import (
	"github.com/amirhnajafiz/authX/internal/model"

	"gorm.io/gorm"
)

type Users interface {
	Insert(user *model.User) error
	GetByID(id uint) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
}

func New(db *gorm.DB) Users {
	return &users{
		db: db,
	}
}

type users struct {
	db *gorm.DB
}

func (u *users) Insert(user *model.User) error {
	return nil
}

func (u *users) GetByID(id uint) (*model.User, error) {
	return nil, nil
}

func (u *users) GetByEmail(email string) (*model.User, error) {
	return nil, nil
}
