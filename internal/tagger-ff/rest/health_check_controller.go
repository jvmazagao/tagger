package rest

import (
	"tagger-ff/internal/tagger-ff/service"

	"github.com/gofiber/fiber/v2"
)

func IsAlive(c *fiber.Ctx) error {
	values, _ := service.IsAlive()
	return c.Status(200).JSON(values)
}
