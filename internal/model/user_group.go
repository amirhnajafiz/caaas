package model

import "time"

// UserGroup reperents a relationship between each user and a group.
type UserGroup struct {
	UserID    int64     `pg:"user_id,unique:usergroup"`
	GroupName string    `pg:"group_name,unique:usergroup"`
	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
}
