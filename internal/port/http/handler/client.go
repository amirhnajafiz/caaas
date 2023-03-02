package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/amirhnajafiz/authX/internal/model"

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

	client := model.Client{
		AppKey:      ctx.Params("app_key"),
		Credentials: claimsString,
	}

	if err := h.Repository.Clients.Create(&client); err != nil {
		return fiber.ErrInternalServerError
	}

	return ctx.SendStatus(http.StatusCreated)
}

// GetAppClient credentials.
func (h *Handler) GetAppClient(ctx *fiber.Ctx) error {
	appID := ctx.Params("app_key")

	id, _ := strconv.Atoi(ctx.Params("client_id"))
	clientID := uint(id)

	client, err := h.Repository.Clients.GetSingle(clientID)
	if err != nil {
		return fiber.ErrNotFound
	}

	if client.AppKey != appID {
		return fiber.ErrForbidden
	}

	return ctx.JSON(client)
}
