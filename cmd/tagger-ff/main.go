package main

import (
	"tagger-ff/cmd"
	"tagger-ff/internal/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db, _ := database.Connect()
	database.Migrate(db)

	app := fiber.New()
	cmd.Routes(app)
	app.Listen(":3000")
}
