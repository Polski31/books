package api

import "github.com/gofiber/fiber/v2"

func setupRoutes(app *fiber.App) {
	app.Get("/getBooks", getBooks)
	app.Post("/addBook", addBook)
	app.Get("/getBook/:isbn", getBook)

	app.Get("/makeCoffee", makeCoffee)
}
