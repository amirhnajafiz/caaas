package v2

import "time"

type UserResponse struct {
	Username  string    `json:"username"`
	Groups    []string  `json:"groups"`
	Roles     []string  `json:"roles"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
