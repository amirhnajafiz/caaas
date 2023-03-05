package model

import (
	"gorm.io/gorm"
)

// Client
// each app can have one or more
// clients, but each client belongs
// to only one app.
type Client struct {
	gorm.Model
	AppKey      string `json:"app_key"`
	ClientID    string `json:"client_id"`
	Credentials string `json:"credentials"`
}
