package controller

import (
	"github.com/amirhnajafiz/caaas/internal/model"
)

func (c *Controller) NewUser(username, password string) error {
	_, err := c.database.Model(&model.User{
		Username: username,
		Password: password,
	}).Insert()

	return err
}
