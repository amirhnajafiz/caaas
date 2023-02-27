package model

import (
	"time"
)

type Client struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}
