package model

import (
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	AppID       uint   `json:"app_id"`
	Credentials string `json:"credentials"`
}
