package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// Authenticate user requests.
func (m *Middleware) Authenticate(ctx *fiber.Ctx) error {
	return ctx.Next()
}
