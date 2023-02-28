package cmd

import (
	"fmt"
	"log"

	"github.com/amirhnajafiz/authX/internal/port/http/handler"
	"github.com/amirhnajafiz/authX/internal/port/http/middleware"
	"github.com/amirhnajafiz/authX/internal/repository"
	"github.com/amirhnajafiz/authX/internal/storage"

	"github.com/gofiber/fiber/v2"
)

// HTTP command.
type HTTP struct{}

// main function of HTTP command.
func (h HTTP) main(port int) {
	app := fiber.New()

	// open db connection
	db, err := storage.NewConnection(storage.Config{})
	if err != nil {
		log.Println(err)

		return
	}

	// create repository
	r := repository.New(db)

	handlerInstance := handler.Handler{
		Repository: r,
	}
	middlewareInstance := middleware.Middleware{
		Repository: r,
	}

	app.Get("/login", handlerInstance.LoginView)
	app.Get("/signup", handlerInstance.SignupView)

	app.Post("/api/login", handlerInstance.Login)
	app.Put("/api/signup", handlerInstance.Signup)

	v1 := app.Use(middlewareInstance.Authenticate)

	v1.Get("/home", handlerInstance.HomeView)

	v1.Put("/api/app", handlerInstance.CreateApp)
	v1.Get("/api/app/:app_id", handlerInstance.GetAppClient)
	v1.Put("/api/app/:app_id/client", handlerInstance.AddClient)
	v1.Get("/api/app/:app_id/client/:client_id", handlerInstance.GetAppClient)

	if err := app.Listen(fmt.Sprintf(":%d", port)); err != nil {
		log.Println(err)
	}
}
