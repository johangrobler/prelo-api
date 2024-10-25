package main

import (
	"prelo/database"
	"prelo/models"
	"prelo/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	database.Connect()

	models.DBMigration()

	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	router.SetupRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("We are alive!") // => 404 "Not Found"
	})
	// handle unavailable route
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
	app.Listen(":8080")

}
