package handler

import (
	"net/http"

	"github.com/amirhnajafiz/authX/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// AddClient to an app.
func (h *Handler) AddClient(ctx *fiber.Ctx) error {
	// getting app key from params
	appKey := ctx.Params("app_key")

	// get client credentials
	credentials := string(ctx.Body())

	client := model.Client{
		AppKey:      appKey,
		ClientID:    uuid.NewString()[:10],
		Credentials: credentials,
	}

	// creating a new client
	if err := h.Repository.Clients.Create(&client); err != nil {
		h.Logger.Error("failed to create client", zap.Error(err))

		return fiber.ErrInternalServerError
	}

	// generate a new jwt token
	token, err := h.Auth.GenerateJWT(client.ClientID, client.AppKey)
	if err != nil {
		h.Logger.Error("failed to generate jwt token", zap.Error(err))

		return fiber.ErrInternalServerError
	}

	return ctx.Status(http.StatusOK).SendString(token)
}
