package router

import (
	"prelo/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/api/v1/brands", handlers.GetBrands)
	app.Post("/api/v1/brands", handlers.CreateBrand)
}
