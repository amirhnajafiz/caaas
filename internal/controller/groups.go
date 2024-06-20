package controller

import (
	"github.com/amirhnajafiz/caaas/internal/model"
)

// NewUserGroup creates a new UserGroup model with given username and group.
func (c *Controller) NewUserGroup(username, groupName string) error {
	_, err := c.database.Model(&model.UserGroup{
		Username:  username,
		GroupName: groupName,
	}).Insert()

	return err
}
