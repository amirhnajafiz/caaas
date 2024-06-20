package controller

import (
	"github.com/amirhnajafiz/caaas/internal/model"
)

// NewUser creates a new User module with given username and password.
func (c *Controller) NewUser(username, password string) error {
	_, err := c.database.Model(&model.User{
		Username: username,
		Password: password,
	}).Insert()

	return err
}

// DeleteUser removes an existing user with its groups.
func (c *Controller) DeleteUser(username string) error {
	tx, err := c.database.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Model(&model.User{}).Where("username = ?", username).Delete(); err != nil {
		return tx.Rollback()
	}

	if _, err := tx.Model(&model.UserGroup{}).Where("username = ?", username).Delete(); err != nil {
		return tx.Rollback()
	}

	return tx.Commit()
}
