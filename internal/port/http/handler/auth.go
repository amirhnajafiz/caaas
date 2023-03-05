package handler

import (
	"net/http"

	"github.com/amirhnajafiz/authX/internal/model"
	"github.com/amirhnajafiz/authX/internal/port/http/request"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// Register a new user into AuthX.
func (h *Handler) Register(ctx *fiber.Ctx) error {
	// get user request
	userRequest := new(request.Register)

	// parse user request
	if err := ctx.BodyParser(&userRequest); err != nil {
		h.Logger.Error("parsing body failed", zap.Error(err))

		return fiber.ErrBadRequest
	}

	// check if user exists or not
	if user, err := h.Repository.Users.Get(userRequest.StudentNumber); err != nil || user == nil {
		h.Logger.Info("user exists", zap.String("student number", userRequest.StudentNumber))

		// check user password
		if userRequest.Password != user.Password {
			h.Logger.Info("wrong password")

			return fiber.ErrUnauthorized
		}

		// get user app
		app, er := h.Repository.Apps.Get(user.ID)
		if er != nil {
			h.Logger.Info("recreating a new app for user")

			// creating a new app for user
			key, e := h.createNewApp(user.ID)
			if e != nil {
				h.Logger.Error("failed to create new app", zap.Error(err))

				return fiber.ErrInternalServerError
			}

			return ctx.Status(http.StatusOK).SendString(key)
		}

		return ctx.Status(http.StatusOK).SendString(app.AppKey)
	}

	user := model.User{
		StudentNumber: userRequest.StudentNumber,
		Password:      userRequest.Password,
	}

	// save user into database
	if err := h.Repository.Users.Create(&user); err != nil {
		h.Logger.Error("failed to create user", zap.Error(err))

		return fiber.ErrInternalServerError
	}

	// create a new app for user
	key, err := h.createNewApp(user.ID)
	if err != nil {
		h.Logger.Error("failed to create new app", zap.Error(err))

		return fiber.ErrInternalServerError
	}

	return ctx.Status(http.StatusOK).SendString(key)
}
