package handler

import (
	"go.uber.org/zap"
	"net/http"
	"time"

	"github.com/amirhnajafiz/authX/internal/model"
	"github.com/amirhnajafiz/authX/internal/port/http/request"

	"github.com/gofiber/fiber/v2"
)

// Signup a new user into AuthX.
func (h *Handler) Signup(ctx *fiber.Ctx) error {
	userRequest := new(request.Register)

	if err := ctx.BodyParser(&userRequest); err != nil {
		h.Logger.Error("parsing body failed", zap.Error(err))

		return fiber.ErrBadRequest
	}

	if user, err := h.Repository.Users.GetByEmail(userRequest.Email); err != nil || user == nil {
		h.Logger.Info("user exists", zap.String("email", userRequest.Email))

		return fiber.ErrNotAcceptable
	}

	user := model.User{
		Email:     userRequest.Email,
		Password:  userRequest.Password,
		CreatedAt: time.Now(),
	}

	if err := h.Repository.Users.Insert(&user); err != nil {
		h.Logger.Error("insert user failed", zap.Error(err))

		return fiber.ErrInternalServerError
	}

	return ctx.SendStatus(http.StatusCreated)
}

// Login a user into AuthX.
func (h *Handler) Login(ctx *fiber.Ctx) error {
	userRequest := new(request.Register)

	if err := ctx.BodyParser(&userRequest); err != nil {
		h.Logger.Error("failed to parse body", zap.Error(err))

		return fiber.ErrBadRequest
	}

	if user, err := h.Repository.Users.GetByEmail(userRequest.Email); err != nil {
		h.Logger.Error("user not found", zap.String("email", userRequest.Email))

		return fiber.ErrNotFound
	} else if user != nil {
		if user.Password != userRequest.Password {
			h.Logger.Info("incorrect password")

			return fiber.ErrNotFound
		}
	} else {
		h.Logger.Error("failed to fined user")

		return fiber.ErrInternalServerError
	}

	return ctx.SendString("token")
}
