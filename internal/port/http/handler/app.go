package handler

import (
	"github.com/amirhnajafiz/authX/internal/model"

	"github.com/google/uuid"
)

// createNewApp function generates a new app.
func (h *Handler) createNewApp(id uint) (string, error) {
	app := model.App{
		AppKey: uuid.NewString()[:15],
		UserID: id,
	}

	// save user app into database
	if err := h.Repository.Apps.Create(&app); err != nil {
		return "", err
	}

	return app.AppKey, nil
}
