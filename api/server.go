package api

import (
	"books/model"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func StartServer() {
	model.ConnectToDatabase()
	app := fiber.New()
	setupRoutes(app)
	err := app.Listen(strings.Join([]string{URL, port}, ":"))
	if err != nil {
		return
	}
}
