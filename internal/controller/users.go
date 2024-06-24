package controller

import (
	"time"

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

// GetUsers returns the list of systems users, it can also filter users
// by a given keyword.
func (c *Controller) GetUsers(keyword string) ([]model.User, error) {
	var users []model.User

	tx := c.database.Model(&users)

	if len(keyword) > 0 {
		tx.Where("username LIKE ?", "%"+keyword+"%")
	}

	if err := tx.Select(); err != nil {
		return users, err
	}

	return users, nil
}

// GetUser by its username.
func (c *Controller) GetUser(username string) (model.User, error) {
	user := new(model.User)

	if err := c.database.Model(user).Where("username = ?", username).Select(); err != nil {
		return *user, err
	}

	return *user, nil
}

// UpdateUser updates the give user's password.
func (c *Controller) UpdateUser(username, password string) error {
	_, err := c.database.Model((*model.User)(nil)).Set("password = ?", password).Set("updated_at = ?", time.Now()).Where("username = ?", username).Update()

	return err
}

// DeleteUser removes an existing user with its groups.
func (c *Controller) DeleteUser(username string) error {
	tx, err := c.database.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Model((*model.User)(nil)).Where("username = ?", username).Delete(); err != nil {
		return tx.Rollback()
	}

	if _, err := tx.Model((*model.UserGroup)(nil)).Where("username = ?", username).Delete(); err != nil {
		return tx.Rollback()
	}

	return tx.Commit()
}
