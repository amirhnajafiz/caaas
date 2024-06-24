package gateway

import (
	"net/http"
	"time"

	"github.com/amirhnajafiz/caaas/pkg/hashing"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (h Handler) login(c echo.Context) error {
	// get user request
	start := time.Now()
	req := new(LoginRequest)

	if err := c.Bind(req); err != nil {
		h.Metrics.AddFailedCall(c.Path())

		return echo.ErrBadRequest
	}

	req.Password = hashing.MD5Hash(req.Password)

	// fetch user
	user, err := h.Ctl.GetUser(req.Username)
	if err != nil {
		h.Metrics.AddFailedCall(c.Path())

		h.Logger.Error("failed to fetch user", zap.String("username", req.Username), zap.Error(err))

		return echo.ErrNotFound
	}

	// check user password match
	if user.Password != req.Password {
		h.Metrics.AddFailedCall(c.Path())

		return echo.ErrForbidden
	}

	// create jwt token
	token, err := h.Auth.GenerateJWT(req.Username)
	if err != nil {
		h.Metrics.AddFailedCall(c.Path())

		h.Logger.Error("failed to create token", zap.Error(err))

		return echo.ErrInternalServerError
	}

	h.Metrics.ObserveLatency(c.Path(), float64(time.Since(start).Milliseconds()))

	return c.String(http.StatusOK, token)
}

func (h Handler) validate(c echo.Context) error {
	return c.JSON(http.StatusOK, ClaimResponse{
		Username: c.Get("username").(string),
	})
}

func (h Handler) groups(c echo.Context) error {
	start := time.Now()
	username := c.Get("username").(string)

	// fetch user groups
	groups, err := h.Ctl.GetUserGroups(username)
	if err != nil {
		h.Metrics.AddFailedCall(c.Path())

		h.Logger.Error("failed to fetch groups", zap.String("username", username), zap.Error(err))

		return echo.ErrNotFound
	}

	h.Metrics.ObserveLatency(c.Path(), float64(time.Since(start).Milliseconds()))

	return c.JSON(http.StatusOK, GroupsResponse{
		Username: username,
		Groups:   groups,
	})
}

func (h Handler) roles(c echo.Context) error {
	start := time.Now()
	username := c.Get("username").(string)

	// fetch user groups
	roles, err := h.Ctl.GetUserRoles(username)
	if err != nil {
		h.Metrics.AddFailedCall(c.Path())

		h.Logger.Error("failed to fetch roles", zap.String("username", username), zap.Error(err))

		return echo.ErrNotFound
	}

	h.Metrics.ObserveLatency(c.Path(), float64(time.Since(start).Milliseconds()))

	return c.JSON(http.StatusOK, RolesResponse{
		Username: username,
		Roles:    roles,
	})
}
