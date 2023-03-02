package model

import (
	"gorm.io/gorm"
)

type App struct {
	gorm.Model
	Name    string `json:"name"`
	UserKey string `json:"user_key"`
	AppKey  string `json:"app_key"`
	UserID  uint   `json:"user_id"`
}
