package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const PORT = "8000"

func main() {
	_, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:33066)/ambassador"), &gorm.Config{})
	if err != nil {
		panic("Could not connect with database")
	}
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":" + PORT)
}
