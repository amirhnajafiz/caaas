package gateway

import (
	"net/http"

	"github.com/amirhnajafiz/caaas/pkg/hashing"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (h Handler) login(c echo.Context) error {
	// get user request
	req := new(LoginRequest)

	if err := c.Bind(req); err != nil {
		return echo.ErrBadRequest
	}

	req.Password = hashing.MD5Hash(req.Password)

	// fetch user
	user, err := h.Ctl.GetUser(req.Username)
	if err != nil {
		h.Logger.Error("failed to fetch user", zap.String("username", req.Username), zap.Error(err))

		return echo.ErrNotFound
	}

	// check user password match
	if user.Password != req.Password {
		return echo.ErrForbidden
	}

	// create jwt token
	token, err := h.Auth.GenerateJWT(req.Username)
	if err != nil {
		h.Logger.Error("failed to create token", zap.Error(err))

		return echo.ErrInternalServerError
	}

	return c.String(http.StatusOK, token)
}

func (h Handler) validate(c echo.Context) error {
	return nil
}

func (h Handler) groups(c echo.Context) error {
	return nil
}
