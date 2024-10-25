package handlers

import (
	"prelo/database"
	"prelo/models"

	"github.com/gofiber/fiber/v2"
)

func GetBrands(c *fiber.Ctx) error {

	db := database.DB.Db
	var brands []models.Brand
	db.Find(&brands)
	return c.JSON(brands)

}

func CreateBrand(c *fiber.Ctx) error {
	db := database.DB.Db
	var brand models.Brand
	if err := c.BodyParser(&brand); err != nil {
		return err
	}
	db.Create(&brand)
	return c.JSON(brand)
}
