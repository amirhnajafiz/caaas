package handler

import "github.com/gofiber/fiber/v2"

func (h *Handler) HomeView(ctx *fiber.Ctx) error {
	return ctx.Render("index", nil)
}

func (h *Handler) RegisterView(ctx *fiber.Ctx) error {
	return ctx.Render("register", nil)
}

func (h *Handler) DocsView(ctx *fiber.Ctx) error {
	return ctx.Render("documents", nil)
}
