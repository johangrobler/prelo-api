package handlers

import (
	"fmt"
	"prelo/database"
	"prelo/models"

	"github.com/gofiber/fiber/v2"
)

func GetItems(c *fiber.Ctx) error {

	userID := c.Locals("userID").(string)
	db := database.DB.Db
	var items []models.Item
	db.Order("CreatedAt desc").Where("user_id = ?", userID).Find(&items)
	return c.JSON(items)

}

func CreateItem(c *fiber.Ctx) error {
	db := database.DB.Db

	userID := c.Locals("userID").(string)
	var item models.Item
	if err := c.BodyParser(&item); err != nil {
		return err
	}
	//item.UserID = userUUID
	fmt.Println("USER ID:", userID)

	item.UserID = userID

	err := db.Create(&item).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create card", "data": err})
	}
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
