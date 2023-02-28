package middleware

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Authenticate user requests.
func (m *Middleware) Authenticate(ctx *fiber.Ctx) error {
	if token := ctx.Get("token", ""); token != "" {
		if email, err := m.Auth.ParseJWT(token); err == nil {
			// set variables into request context
			ctx.Locals("email", email)

			return ctx.Next()
		} else {
			log.Println(err)

			// 401
			return ctx.SendStatus(http.StatusUnauthorized)
		}
	} else {
		// 400
		return ctx.SendStatus(http.StatusBadRequest)
	}
}
