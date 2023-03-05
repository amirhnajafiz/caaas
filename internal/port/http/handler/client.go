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

	return ctx.Status(http.StatusOK).SendString(client.ClientID)
}

// CheckClient checks a existent client.
func (h *Handler) CheckClient(ctx *fiber.Ctx) error {
	// getting app key from params
	appKey := ctx.Params("app_key")

	// get user token
	clientID := ctx.Query("id")

	// get client from database
	client, err := h.Repository.Clients.Get(clientID)
	if err != nil {
		h.Logger.Info("failed to find client", zap.Error(err))

		return fiber.ErrNotFound
	}

	if client.AppKey != appKey {
		return fiber.ErrUnauthorized
	}

	// generate a new jwt token
	token, err := h.Auth.GenerateJWT(client.ClientID, client.AppKey)
	if err != nil {
		h.Logger.Error("failed to generate jwt token", zap.Error(err))

		return fiber.ErrInternalServerError
	}

	return ctx.Status(http.StatusOK).SendString(token)
}

// GetClient by token.
func (h *Handler) GetClient(ctx *fiber.Ctx) error {
	// getting app key from params
	appKey := ctx.Params("app_key")

	// get user token
	token := ctx.Query("token")

	clientID, key, err := h.Auth.ParseJWT(token)
	if err != nil {
		h.Logger.Error("parse token failed", zap.Error(err))

		return fiber.ErrBadRequest
	}

	// check app keys
	if key != appKey {
		return fiber.ErrUnauthorized
	}

	// get client from database
	client, err := h.Repository.Clients.Get(clientID)
	if err != nil {
		h.Logger.Error("failed to find user", zap.Error(err))

		return fiber.ErrNotFound
	}

	return ctx.Status(http.StatusOK).SendString(client.Credentials)
}
