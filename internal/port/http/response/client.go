package response

import "time"

type Client struct {
	Claims    string    `json:"claims"`
	CreatedAt time.Time `json:"created_at"`
}
