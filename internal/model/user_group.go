package model

import "time"

// UserGroup reperents a relationship between each user and a group.
type UserGroup struct {
	Username  string    `pg:"username,unique:usergroup"`
	GroupName string    `pg:"group_name,unique:usergroup"`
	CreatedAt time.Time `pg:"created_at,default:now()"`
	UpdatedAt time.Time `pg:"updated_at,default:now()"`
}
