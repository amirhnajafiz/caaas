package handler

import (
	"github.com/gofiber/fiber/v2"
)

// RootView of web application.
func (h *Handler) RootView(ctx *fiber.Ctx) error {
	return ctx.Render("root", fiber.Map{
		"uri": ctx.Hostname(),
	})
}

// LoginView returns the login view.
func (h *Handler) LoginView(ctx *fiber.Ctx) error {
	return ctx.Render("login", fiber.Map{
		"uri": ctx.Hostname(),
	})
}

// SignupView returns the signup view.
func (h *Handler) SignupView(ctx *fiber.Ctx) error {
	return ctx.Render("signup", fiber.Map{
		"uri": ctx.Hostname(),
	})
}

// HomeView returns the home page of application.
func (h *Handler) HomeView(ctx *fiber.Ctx) error {
	// todo: get user apps
	return ctx.Render("home", fiber.Map{
		"uri": ctx.Hostname(),
	})
}
