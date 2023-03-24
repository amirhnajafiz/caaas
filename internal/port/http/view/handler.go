package view

import "github.com/gofiber/fiber/v2"

type View struct{}

func (v *View) HomeView(ctx *fiber.Ctx) error {
	return ctx.Render("index", nil)
}

func (v *View) RegisterView(ctx *fiber.Ctx) error {
	return ctx.Render("register", nil)
}

func (v *View) DocsView(ctx *fiber.Ctx) error {
	return ctx.Render("documents", nil)
}
