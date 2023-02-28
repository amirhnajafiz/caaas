package handler

import (
	"strconv"
	"time"

	"github.com/amirhnajafiz/authX/internal/model"
	"github.com/amirhnajafiz/authX/internal/port/http/request"
	"github.com/amirhnajafiz/authX/internal/port/http/response"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// CreateApp for a user.
func (h *Handler) CreateApp(ctx *fiber.Ctx) error {
	userRequest := new(request.NewApp)

	if err := ctx.BodyParser(&userRequest); err != nil {
		return fiber.ErrBadRequest
	}

	appInstance := model.App{
		Name:      userRequest.Name,
		Key:       uuid.New().String()[:15],
		URI:       uuid.NewString()[:10],
		UserID:    0, // todo: find user id from token
		CreatedAt: time.Now(),
	}

	if err := h.Repository.Apps.Create(&appInstance); err != nil {
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(fiber.Map{
		"api_key": appInstance.Key,
		"uri_id":  appInstance.URI,
	})
}

// GetSingleApp of a user.
func (h *Handler) GetSingleApp(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("app_id"))
	appID := uint(id)

	app, err := h.Repository.Apps.GetSingle(appID)
	if err != nil {
		return fiber.ErrNotFound
	}

	clients, err := h.Repository.Clients.GetAppClients(app.ID)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(response.AppResponse{
		ID:      appID,
		Name:    app.Name,
		Clients: clients,
	})
}