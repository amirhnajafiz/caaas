package handler

import (
	"net/http"

	"github.com/amirhnajafiz/authX/internal/model"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// AddClient to an app.
func (h *Handler) AddClient(ctx *fiber.Ctx) error {
	claims := string(ctx.Body())

	app, err := h.Repository.Apps.GetByKey(ctx.Params("app_key"))
	if err != nil {
		h.Logger.Error("app not found", zap.Error(err))

		return fiber.ErrNotFound
	}

	client := model.Client{
		AppID:       app.ID,
		Credentials: claims,
	}

	if err := h.Repository.Clients.Create(&client); err != nil {
		h.Logger.Error("failed to create client", zap.Error(err))

		return fiber.ErrInternalServerError
	}

	return ctx.SendStatus(http.StatusCreated)
}
