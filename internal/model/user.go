package model

import "time"

// Each user has a unique username and a password.
type User struct {
	Username  string    `pg:"username,unique"`
	Password  string    `pg:"password"`
	CreatedAt time.Time `pg:"created_at,default:now()"`
	UpdatedAt time.Time `pg:"updated_at,default:now()"`
}
