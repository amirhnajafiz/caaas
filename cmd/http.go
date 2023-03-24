package cmd

import (
	"fmt"

	"github.com/amirhnajafiz/authX/internal/config"
	"github.com/amirhnajafiz/authX/internal/port/http/handler"
	"github.com/amirhnajafiz/authX/internal/port/http/view"
	"github.com/amirhnajafiz/authX/internal/repository"
	"github.com/amirhnajafiz/authX/internal/storage"
	"github.com/amirhnajafiz/authX/pkg/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// HTTP command.
type HTTP struct {
	Cfg    config.Config
	Logger *zap.Logger
}

// Command returns the cobra command.
func (h HTTP) Command() *cobra.Command {
	run := func(cmd *cobra.Command, args []string) { h.main() }
	return &cobra.Command{Use: "http", Run: run}
}

// main function of HTTP command.
func (h HTTP) main() {
	// initialize standard Go html template engine
	engine := html.New("./views", ".html")

	// create a new fiber app
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// open db connection
	db, err := storage.NewConnection(h.Cfg.Storage)
	if err != nil {
		h.Logger.Error("database connection failed", zap.Error(err))

		return
	}

	// create a new auth
	a := auth.New(h.Cfg.Auth)

	// create repository
	r := repository.New(db)

	handlerInstance := handler.Handler{
		Auth:       a,
		Logger:     h.Logger.Named("handler"),
		Repository: r,
	}

	viewInstance := view.View{}

	app.Static("/", "./public")

	app.Get("/", viewInstance.Home)
	app.Get("/docs", viewInstance.Docs)
	app.Get("/register", viewInstance.Register)

	app.Post("/api/register", handlerInstance.Register)
	app.Post("/api/app/:app_key", handlerInstance.CheckClient)
	app.Put("/api/app/:app_key", handlerInstance.AddClient)
	app.Get("/api/app/:app_key", handlerInstance.GetClient)

	if er := app.Listen(fmt.Sprintf(":%d", h.Cfg.HTTP.Port)); er != nil {
		h.Logger.Error("failed to start app", zap.Error(er))
	}
}
