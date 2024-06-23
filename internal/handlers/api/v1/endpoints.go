package v1

import (
	"net/http"

	"github.com/amirhnajafiz/caaas/pkg/hashing"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (h Handler) getAllUsers(c echo.Context) error {
	keyword := c.QueryParam("keyword")

	// fetch users
	users, err := h.Ctl.GetUsers(keyword)
	if err != nil {
		h.Logger.Error("failed to get users", zap.Error(err))

		return echo.ErrInternalServerError
	}

	// convert them to response
	list := make([]UserResponse, len(users))
	for index, user := range users {
		list[index] = UserResponse{
			Username:  user.Username,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
	}

	return c.JSON(http.StatusOK, list)
}

func (h Handler) createUser(c echo.Context) error {
	// get user request
	req := new(UserRequest)
	if err := c.Bind(req); err != nil {
		return echo.ErrBadRequest
	}

	req.Password = hashing.MD5Hash(req.Password)

	// create new user
	if err := h.Ctl.NewUser(req.Username, req.Password); err != nil {
		h.Logger.Error("failed to create user", zap.Error(err))
	}

	return c.String(http.StatusOK, "")
}

func (h Handler) updateUser(c echo.Context) error {
	// get user request
	req := new(UserRequest)
	if err := c.Bind(req); err != nil {
		return echo.ErrBadRequest
	}

	req.Password = hashing.MD5Hash(req.Password)

	// update user
	if err := h.Ctl.UpdateUser(req.Username, req.Password); err != nil {
		h.Logger.Error("failed to update user", zap.Error(err))
	}

	return c.String(http.StatusOK, "")
}

func (h Handler) removeUser(c echo.Context) error {
	username := c.QueryParam("username")

	// remove user
	if err := h.Ctl.DeleteUser(username); err != nil {
		h.Logger.Error("failed to remove a user", zap.String("username", username), zap.Error(err))

		return echo.ErrInternalServerError
	}

	return c.String(http.StatusOK, "")
}

func (h Handler) addUserToGroup(c echo.Context) error {
	// fetch query params
	req := new(UserGroupQuery)
	if err := c.Bind(req); err != nil {
		return echo.ErrBadRequest
	}

	// add user to a group
	if err := h.Ctl.NewUserGroup(req.Username, req.Group); err != nil {
		h.Logger.Error("failed to add user to a group", zap.String("username", req.Username), zap.String("group", req.Group), zap.Error(err))

		return echo.ErrInternalServerError
	}

	return c.String(http.StatusOK, "")
}

func (h Handler) removeUserFromGroup(c echo.Context) error {
	// fetch query params
	req := new(UserGroupQuery)
	if err := c.Bind(req); err != nil {
		return echo.ErrBadRequest
	}

	// remove user from a gorup
	if err := h.Ctl.RemoveUserGroup(req.Username, req.Group); err != nil {
		h.Logger.Error("failed to remove user from a group", zap.String("username", req.Username), zap.String("group", req.Group), zap.Error(err))

		return echo.ErrInternalServerError
	}

	return c.String(http.StatusOK, "")
}

func (h Handler) removeGroup(c echo.Context) error {
	group := c.QueryParam("group")

	// remove a group
	if err := h.Ctl.RemoveGroup(group); err != nil {
		h.Logger.Error("failed to remove a group", zap.String("group", group), zap.Error(err))

		return echo.ErrInternalServerError
	}

	return c.String(http.StatusOK, "")
}
