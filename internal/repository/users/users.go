package users

import (
	"fmt"

	"github.com/amirhnajafiz/authX/internal/model"

	"gorm.io/gorm"
)

// Users manages the user model.
type Users interface {
	Create(user *model.User) error
	Get(studentNumber string) (*model.User, error)
}

// New generates a new user repository.
func New(db *gorm.DB) Users {
	return &users{
		db: db,
	}
}

// users manages the functions of repository.
type users struct {
	db *gorm.DB
}

// Create a new user.
func (u *users) Create(user *model.User) error {
	return u.db.Create(user).Error
}

// Get a user by its student number.
func (u *users) Get(studentNumber string) (*model.User, error) {
	user := new(model.User)

	if err := u.db.Where("student_number = ?", studentNumber).Find(&user).Error; err != nil {
		return nil, err
	}

	// user not found
	if user.StudentNumber != studentNumber {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}
