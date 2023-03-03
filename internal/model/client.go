package model

import (
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	AppID       string `json:"app_id"`
	Credentials string `json:"credentials"`
}
