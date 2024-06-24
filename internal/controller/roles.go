package controller

import (
	"github.com/amirhnajafiz/caaas/internal/model"
	"github.com/go-pg/pg/v10"
)

// NewUserRole creates a new UserRole model with given username and role.
func (c *Controller) NewUserRole(username, role string) error {
	_, err := c.database.Model(&model.UserRole{
		Username: username,
		Role:     role,
	}).Insert()

	return err
}

// GetUserRoles returns a list of user roles.
func (c *Controller) GetUserRoles(username string) ([]string, error) {
	var roles []string

	if err := c.database.Model((*model.UserRole)(nil)).ColumnExpr("array_agg(role)").Where("username = ?", username).Select(pg.Array(&roles)); err != nil {
		return roles, err
	}

	return roles, nil
}

// RemoveRole removes all user role records with given role.
func (c *Controller) RemoveRole(role string) error {
	_, err := c.database.Model((*model.UserRole)(nil)).Where("role = ?", role).Delete()

	return err
}

// RemoveUserRole removes a user role record.
func (c *Controller) RemoveUserRole(username, role string) error {
	_, err := c.database.Model((*model.UserRole)(nil)).Where("username = ? AND role = ?", username, role).Delete()

	return err
}
