package cmd

import (
	"fmt"

	"github.com/amirhnajafiz/authX/internal/config"
	"github.com/amirhnajafiz/authX/internal/port/http/handler"
	"github.com/amirhnajafiz/authX/internal/port/http/middleware"
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
	// create fiber engine
	engine := html.New("./views", ".html")

	// create a new fiber app
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts",
	})

	// open db connection
	db, err := storage.NewConnection(h.Cfg.Storage)
	if err != nil {
		h.Logger.Error("database connection failed", zap.Error(err))

		return
	}

	// create repository
	r := repository.New(db)

	handlerInstance := handler.Handler{
		Logger:     h.Logger.Named("handler"),
		Repository: r,
	}
	middlewareInstance := middleware.Middleware{
		Auth:       *auth.New(h.Cfg.Auth),
		Logger:     h.Logger.Named("middleware"),
		Repository: r,
	}

	app.Get("/", handlerInstance.RootView)
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

	if er := app.Listen(fmt.Sprintf(":%d", h.Cfg.HTTP.Port)); er != nil {
		h.Logger.Error("start app failed", zap.Error(er))
	}
}
