package handlers

import (
	"prelo/database"
	"prelo/models"

	"github.com/gofiber/fiber/v2"
)

func GetCategories(c *fiber.Ctx) error {

	db := database.DB.Db
	var categories []models.Category
	db.Find(&categories)
	return c.JSON(categories)

}

func CreateCategory(c *fiber.Ctx) error {
	db := database.DB.Db
	var category models.Category
	if err := c.BodyParser(&category); err != nil {
		return err
	}
	db.Create(&category)
	return c.JSON(category)
}
