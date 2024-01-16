package main

import (
    "github.com/gofiber/fiber/v2"
	"rawgo.com/m/handlers"
)

func setupRoutes(app *fiber.App) {
    app.Get("/", handlers.Home)
}
