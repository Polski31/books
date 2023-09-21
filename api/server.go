package api

import (
	"books/model"
	"github.com/gofiber/fiber/v2"
	"strings"
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
