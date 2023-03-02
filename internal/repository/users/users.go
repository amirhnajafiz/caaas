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
	return u.db.Create(user).Error
}

func (u *users) GetByID(id uint) (*model.User, error) {
	user := new(model.User)

	if err := u.db.First(user).Where("id = ?", id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *users) GetByEmail(email string) (*model.User, error) {
	user := new(model.User)

	if err := u.db.Find(user).Where("email = ?", email).Error; err != nil {
		return nil, err
	}

	return user, nil
}
