package cmd

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type HTTP struct{}

func (h HTTP) main(port int) {
	app := fiber.New()

	if err := app.Listen(fmt.Sprintf(":%d", port)); err != nil {
		log.Println(err)
	}
}
