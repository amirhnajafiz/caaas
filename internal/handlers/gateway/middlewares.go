package gateway

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// requestsMiddleware runs before each request.
func (h Handler) requestsMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		h.Metrics.AddRequest(c.Path())
		return next(c)
	}
}

// authMiddleware is used to check jwt token
func (h Handler) authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// get token from request headers
		token := c.Request().Header.Get("X-token")

		// if not exists
		if len(token) == 0 {
			h.Logger.Debug("no give token")

			return echo.ErrBadRequest
		}

		// if token was invalid, return 401
		if username, err := h.Auth.ParseJWT(token); err != nil {
			h.Logger.Debug("failed to parse JWT token", zap.String("token", token))

			return echo.ErrUnauthorized
		} else {
			// save username into context
			c.Set("username", username)
		}

		return next(c)
	}
}
