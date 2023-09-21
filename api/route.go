package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get(fmt.Sprintf("%s/%s", basePath, "getBooks"), getBooks)
	app.Post(fmt.Sprintf("%s/%s", basePath, "addBook"), addBook)
	app.Get(fmt.Sprintf("%s/%s/:%s", basePath, "getBook", urlFilterVariable), getBook)

	app.Get(fmt.Sprintf("%s/%s", basePath, "makeCoffee"), makeCoffee)
}
