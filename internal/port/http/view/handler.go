package view

import "github.com/gofiber/fiber/v2"

type View struct{}

func (v *View) Home(ctx *fiber.Ctx) error {
	return ctx.Render("index", nil)
}

func (v *View) Register(ctx *fiber.Ctx) error {
	return ctx.Render("register", nil)
}

func (v *View) Docs(ctx *fiber.Ctx) error {
	return ctx.Render("documents", nil)
}
