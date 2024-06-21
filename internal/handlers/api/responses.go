package api

import "time"

type UserResponse struct {
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserGroupsResponse struct {
	Username string   `json:"username"`
	Groups   []string `json:"groups"`
}