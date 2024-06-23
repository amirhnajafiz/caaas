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

	return c.JSON(http.StatusOK, UserResponse{
		Username:  user.Username,
		Groups:    groups,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
}
