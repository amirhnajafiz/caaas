package v2

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (h Handler) getUser(c echo.Context) error {
	username := c.QueryParam("username")

	// fetch user
	user, err := h.Ctl.GetUser(username)
	if err != nil {
		h.Logger.Error("failed to get user", zap.String("username", username), zap.Error(err))

		return echo.ErrNotFound
	}

	// fetch user's groups
	groups, err := h.Ctl.GetUserGroups(user.Username)
	if err != nil {
		h.Logger.Error("failed to get user's groups", zap.String("username", username), zap.Error(err))

		return echo.ErrInternalServerError
	}

	// fetch user's roles
	roles, err := h.Ctl.GetUserRoles(user.Username)
	if err != nil {
		h.Logger.Error("failed to get user's roles", zap.String("username", username), zap.Error(err))

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, UserResponse{
		Username:  user.Username,
		Groups:    groups,
		Roles:     roles,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
}

func (h Handler) addRoleToUser(c echo.Context) error {
	// fetch query params
	req := new(UserRoleQuery)
	if err := c.Bind(req); err != nil {
		return echo.ErrBadRequest
	}

	// add role to a user
	if err := h.Ctl.NewUserRole(req.Username, req.Role); err != nil {
		h.Logger.Error("failed to add role to a user", zap.String("username", req.Username), zap.String("role", req.Role), zap.Error(err))

		return echo.ErrInternalServerError
	}

	return c.String(http.StatusOK, "")
}

func (h Handler) removeRoleFromUser(c echo.Context) error {
	// fetch query params
	req := new(UserRoleQuery)
	if err := c.Bind(req); err != nil {
		return echo.ErrBadRequest
	}

	// remove role from a user
	if err := h.Ctl.RemoveUserGroup(req.Username, req.Role); err != nil {
		h.Logger.Error("failed to remove role from a user", zap.String("username", req.Username), zap.String("role", req.Role), zap.Error(err))

		return echo.ErrInternalServerError
	}

	return c.String(http.StatusOK, "")
}

func (h Handler) removeRole(c echo.Context) error {
	role := c.QueryParam("role")

	// remove a role
	if err := h.Ctl.RemoveRole(role); err != nil {
		h.Logger.Error("failed to remove a role", zap.String("role", role), zap.Error(err))

		return echo.ErrInternalServerError
	}

	return c.String(http.StatusOK, "")
}
