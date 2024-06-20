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

// RemoveGroup removes all user group records with given group_name.
func (c *Controller) RemoveGroup(groupName string) error {
	_, err := c.database.Model(&model.UserGroup{}).Where("group_name = ?", groupName).Delete()

	return err
}

// RemoveUserGroup removes a user group record.
func (c *Controller) RemoveUserGroup(username, groupName string) error {
	_, err := c.database.Model(&model.UserGroup{}).Where("username = ? AND group_name = ?", username, groupName).Delete()

	return err
}
