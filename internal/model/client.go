package model

import (
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	AppKey      string `json:"app_key"`
	Credentials string `json:"credentials"`
}
