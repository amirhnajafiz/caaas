package users

import (
	"github.com/amirhnajafiz/authX/internal/model"
)

type Users interface {
	Insert(user *model.User) error
	GetByID(id uint) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
}

func New() Users {
	return &users{}
}

type users struct{}

func (u *users) Insert(user *model.User) error {
	return nil
}

func (u *users) GetByID(id uint) (*model.User, error) {
	return nil, nil
}

func (u *users) GetByEmail(email string) (*model.User, error) {
	return nil, nil
}
