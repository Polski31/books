package api

import (
	"github.com/gofiber/fiber/v2"
)

func makeCoffee(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusTeapot)
}
