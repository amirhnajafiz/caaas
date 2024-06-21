package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (h Handler) getAllUsers(c echo.Context) error {
	return nil
}

func (h Handler) createUser(c echo.Context) error {
	return nil
}

func (h Handler) updateUser(c echo.Context) error {
	return nil
}

func (h Handler) removeUser(c echo.Context) error {
	username := c.QueryParam("username")

	if err := h.Ctl.DeleteUser(username); err != nil {
		h.Logger.Error("failed to remove a user", zap.String("username", username), zap.Error(err))

		return echo.ErrInternalServerError
	}

	return c.String(http.StatusOK, "")
}

func (h Handler) addUserToGroup(c echo.Context) error {
	group := c.QueryParam("group")
	username := c.QueryParam("username")

	if err := h.Ctl.NewUserGroup(username, group); err != nil {
		h.Logger.Error("failed to add user to a group", zap.String("username", username), zap.String("group", group), zap.Error(err))

		return echo.ErrInternalServerError
	}

	return c.String(http.StatusOK, "")
}

func (h Handler) removeUserFromGroup(c echo.Context) error {
	group := c.QueryParam("group")
	username := c.QueryParam("username")

	if err := h.Ctl.RemoveUserGroup(username, group); err != nil {
		h.Logger.Error("failed to remove user from a group", zap.String("username", username), zap.String("group", group), zap.Error(err))

		return echo.ErrInternalServerError
	}

	return c.String(http.StatusOK, "")
}

func (h Handler) removeGroup(c echo.Context) error {
	group := c.QueryParam("group")

	if err := h.Ctl.RemoveGroup(group); err != nil {
		h.Logger.Error("failed to remove a group", zap.String("group", group), zap.Error(err))

		return echo.ErrInternalServerError
	}

	return c.String(http.StatusOK, "")
}
