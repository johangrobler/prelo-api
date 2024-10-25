package handlers

import (
	"prelo/database"
	"prelo/models"

	"github.com/gofiber/fiber/v2"
)

func GetItems(c *fiber.Ctx) error {

	db := database.DB.Db
	var items []models.Item
	db.Find(&items)
	return c.JSON(items)

}

func CreateItem(c *fiber.Ctx) error {
	db := database.DB.Db
	userID := c.Locals("userID")

	var item models.Item
	if err := c.BodyParser(&item); err != nil {
		return err
	}
	item.UserID = userID.(string)
	db.Create(&item)
	return c.JSON(item)
}

func UpdateItem(c *fiber.Ctx) error {
	db := database.DB.Db
	var item models.Item
	if err := c.BodyParser(&item); err != nil {
		return err
	}
	db.Save(&item)
	return c.JSON(item)
}

func DeleteItem(c *fiber.Ctx) error {
	db := database.DB.Db
	var item models.Item
	if err := c.BodyParser(&item); err != nil {
		return err
	}
	db.Delete(&item)
	return c.JSON(item)
}

func GetItem(c *fiber.Ctx) error {
	db := database.DB.Db
	var item models.Item
	if err := c.BodyParser(&item); err != nil {
		return err
	}
	db.Find(&item)
	return c.JSON(item)
}

func GetItemsByCategoryID(c *fiber.Ctx) error {
	db := database.DB.Db
	var items []models.Item
	if err := c.BodyParser(&items); err != nil {
		return err
	}
	db.Find(&items)
	return c.JSON(items)
}

func GetItemsByBrandID(c *fiber.Ctx) error {
	db := database.DB.Db
	var items []models.Item
	if err := c.BodyParser(&items); err != nil {
		return err
	}
	db.Find(&items)
	return c.JSON(items)
}

func GetItemsByUserID(c *fiber.Ctx) error {
	db := database.DB.Db
	var items []models.Item
	if err := c.BodyParser(&items); err != nil {
		return err
	}
	db.Find(&items)
	return c.JSON(items)
}
