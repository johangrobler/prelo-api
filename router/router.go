package router

import (
	"prelo/handlers"
	"prelo/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	jwt := middlewares.NewAuthMiddleware()
	api := app.Group("/v1")

	//auth
	auth := api.Group("/auth")
	auth.Post("/login", handlers.Login)
	auth.Post("/register", handlers.RegisterUser)
	auth.Get("/me", jwt, handlers.Protected)

	api.Get("brands", handlers.GetBrands)
	api.Post("brands", handlers.CreateBrand)

	api.Get("categories", handlers.GetCategories)
	api.Post("categories", handlers.CreateCategory)
}
