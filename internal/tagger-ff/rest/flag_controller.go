package rest

import (
	"tagger-ff/internal/tagger-ff/service"
	"tagger-ff/pkg/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllFlags(c *fiber.Ctx) error {
	flags, _ := service.GetAllFlags()

	return c.Status(200).JSON(flags)
}

func CreateFlag(c *fiber.Ctx) error {
	var data models.FlagData
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	flag, _ := service.CreateFlag(&data)

	return c.Status(200).JSON(flag)
}
