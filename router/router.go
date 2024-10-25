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

	//items

	api.Get("items", handlers.GetItems)
	api.Post("items", jwt, handlers.CreateItem)
	api.Get("items/:id", handlers.GetItem)
	api.Put("items/:id", jwt, handlers.UpdateItem)
	api.Delete("items/:id", jwt, handlers.DeleteItem)
	api.Get("items/user/:id", handlers.GetItemsByUserID)
	api.Get("items/brand/:id", handlers.GetItemsByBrandID)
	api.Get("items/category/:id", handlers.GetItemsByCategoryID)

}
