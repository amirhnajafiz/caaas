package model

import "time"

type App struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Key       string    `json:"key"`
	URI       string    `json:"uri"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
