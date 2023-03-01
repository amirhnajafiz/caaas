package handler

import (
	"fmt"
	"github.com/amirhnajafiz/authX/internal/model"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// AddClient to an app.
func (h *Handler) AddClient(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("app_id"))
	appID := uint(id)

	claims := make(map[string]interface{})
	if err := ctx.BodyParser(&claims); err != nil {
		return fiber.ErrBadRequest
	}

	claimsString := ""
	for key := range claims {
		claimsString = fmt.Sprintf("%s&%s=%s", claimsString, key, claims[key])
	}

	client := model.Client{
		AppID:       appID,
		Credentials: claimsString,
	}

	if err := h.Repository.Clients.Create(&client); err != nil {
		return fiber.ErrInternalServerError
	}

	return ctx.SendStatus(http.StatusCreated)
}

// GetAppClient credentials.
func (h *Handler) GetAppClient(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("app_id"))
	appID := uint(id)

	id, _ = strconv.Atoi(ctx.Params("client_id"))
	clientID := uint(id)

	return nil
}
