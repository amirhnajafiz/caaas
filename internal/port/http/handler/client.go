package handler

import (
	"fmt"
	"github.com/amirhnajafiz/authX/internal/model"
	"go.uber.org/zap"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// AddClient to an app.
func (h *Handler) AddClient(ctx *fiber.Ctx) error {
	claims := make(map[string]interface{})
	if err := ctx.BodyParser(&claims); err != nil {
		return fiber.ErrBadRequest
	}

	claimsString := ""
	for key := range claims {
		claimsString = fmt.Sprintf("%s&%s=%s", claimsString, key, claims[key])
	}

	app, err := h.Repository.Apps.GetByKey(ctx.Params("app_key"))
	if err != nil {
		h.Logger.Error("app not found", zap.Error(err))

		return fiber.ErrNotFound
	}

	client := model.Client{
		AppID:       app.ID,
		Credentials: claimsString,
	}

	if err := h.Repository.Clients.Create(&client); err != nil {
		return fiber.ErrInternalServerError
	}

	return ctx.SendStatus(http.StatusCreated)
}
