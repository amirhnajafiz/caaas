package model

import (
	"gorm.io/gorm"
)

// App
// each user has an app that we create it
// when user is registering.
type App struct {
	gorm.Model
	AppKey string `json:"app_key"`
	UserID uint   `json:"user_id"`
}
