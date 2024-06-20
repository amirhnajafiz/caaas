package controller

import (
	"github.com/amirhnajafiz/caaas/internal/model"
)

func (c *Controller) NewUserGroup(username, groupName string) error {
	_, err := c.database.Model(&model.UserGroup{
		Username:  username,
		GroupName: groupName,
	}).Insert()

	return err
}
