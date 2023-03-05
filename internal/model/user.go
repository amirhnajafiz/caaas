package model

import (
	"gorm.io/gorm"
)

// User
// each user is a student that we store
// their student number and a password.
type User struct {
	gorm.Model
	StudentNumber string `json:"student_number"`
	Password      string `json:"password"`
}
