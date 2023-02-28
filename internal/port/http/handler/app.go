package handler

import (
	"time"

	"github.com/amirhnajafiz/authX/internal/model"
	"github.com/amirhnajafiz/authX/internal/port/http/request"

	"github.com/gofiber/fiber/v2"
)

// CreateApp for a user.
func (h *Handler) CreateApp(ctx *fiber.Ctx) error {
	userRequest := new(request.NewApp)

	if err := ctx.BodyParser(&userRequest); err != nil {
		return fiber.ErrBadRequest
	}

	appInstance := model.App{
		Name:      userRequest.Name,
		Key:       "", // todo: create a api key
		URI:       "", // todo: create a unique id
		UserID:    0,  // todo: find user id from token
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
	return nil
}

// RemoveApp of a user.
func (h *Handler) RemoveApp(ctx *fiber.Ctx) error {
	return nil
}
