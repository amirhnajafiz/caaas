package model

import "time"

// UserRole reperents a relationship between each user and a role.
type UserRole struct {
	Username  string    `pg:"username,unique:userrole"`
	Role      string    `pg:"role,unique:userrole"`
	CreatedAt time.Time `pg:"created_at,default:now()"`
	UpdatedAt time.Time `pg:"updated_at,default:now()"`
}
