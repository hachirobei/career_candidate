package main

import "github.com/gofiber/fiber/v2"
import "rawgo.com/m/database"

func main() {

    database.ConnectDb()

    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, Muhammad Hafizoddin!")
    })

    app.Listen(":1000")
}