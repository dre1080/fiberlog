package main

import (
	"github.com/dre1080/fiberlog"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	// Default
	app.Use(fiberlog.New())

	app.Get("/ok", func(c *fiber.Ctx) {
		c.SendString("ok")
	})

	app.Get("/warn", func(c *fiber.Ctx) {
		c.Next(fiber.ErrUnprocessableEntity)
	})

	app.Get("/err", func(c *fiber.Ctx) {
		c.Next(fiber.ErrInternalServerError)
	})

	app.Listen(3000)
}
