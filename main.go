package main

import (
	"FiberProject/src/database"
	"FiberProject/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

const PORT = "8000"

func main() {
	database.Connect()
	database.AutoMigrate()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		//AllowOrigins: "*",
	}))
	routes.Setup(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World🎇")
	})

	app.Listen(":" + PORT)
}
