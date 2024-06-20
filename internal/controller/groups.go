package controller

import (
	"github.com/amirhnajafiz/caaas/internal/model"

	"github.com/go-pg/pg/v10"
)

// NewUserGroup creates a new UserGroup model with given username and group.
func (c *Controller) NewUserGroup(username, groupName string) error {
	_, err := c.database.Model(&model.UserGroup{
		Username:  username,
		GroupName: groupName,
	}).Insert()

	return err
}

// GetUserGroups returns a list of user groups.
func (c *Controller) GetUserGroups(username string) ([]string, error) {
	var groups []string

	if err := c.database.Model((*model.UserGroup)(nil)).ColumnExpr("array_agg(group_name)").Where("username = ?", username).Select(pg.Array(&groups)); err != nil {
		return groups, err
	}

	return groups, nil
}

// RemoveGroup removes all user group records with given group_name.
func (c *Controller) RemoveGroup(groupName string) error {
	_, err := c.database.Model((*model.UserGroup)(nil)).Where("group_name = ?", groupName).Delete()

	return err
}

// RemoveUserGroup removes a user group record.
func (c *Controller) RemoveUserGroup(username, groupName string) error {
	_, err := c.database.Model((*model.UserGroup)(nil)).Where("username = ? AND group_name = ?", username, groupName).Delete()

	return err
}
