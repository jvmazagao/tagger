package cmd

import (
	"tagger-ff/internal/tagger-ff/rest"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/IsAlive", rest.IsAlive)

	// Flags
	app.Get("/flags", rest.GetAllFlags)
}
