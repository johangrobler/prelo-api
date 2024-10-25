package router

import (
	"prelo/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/v1")

	api.Get("brands", handlers.GetBrands)
	api.Post("brands", handlers.CreateBrand)

	api.Get("categories", handlers.GetCategories)
	api.Post("categories", handlers.CreateCategory)
}
