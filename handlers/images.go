package handlers

import (
	"fmt"
	"prelo/database"
	"prelo/models"
	"prelo/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// UploadFile -
func UploadFile(c *fiber.Ctx) error {
	fmt.Println("Upload route!")
	file, err := c.FormFile("file")

	fmt.Println("uploadimg!!!", err)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	fmt.Println("saved!!!")
	// Get Buffer from file\

	buffer, err := file.Open()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	defer buffer.Close()

	objectName := file.Filename

	contentType := file.Header["Content-Type"][0]
	fileSize := file.Size

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	db := database.DB.Db

	userID := c.Locals("userID").(string)
	fmt.Println("uid", userID)
	id := c.Params("id")

	image := new(models.Image)
	image.ID = uuid.New()
	image.UserID = userID
	image.ItemID = id
	image.Path = fmt.Sprintf("uploads/%s-%s", image.ID.String(), file.Filename)
	image.Position = 1

	c.SaveFile(file, image.Path)

	//services.UploadImages(image.Path)
	//fmt.Println(image.ID.String() + "-" + file.Filename)
	//image.Url = services.UploadToSpaces("properties", fmt.Sprintf("%s-%s", image.ID.String(), file.Filename))
	id = userID
	services.MakeImages(image, id)
	db.Create(&image)

	return c.JSON(fiber.Map{
		"error":       false,
		"msg":         nil,
		"objectName":  objectName,
		"contentType": contentType,
		"fileSize":    fileSize,
		"url":         image.Path,
	})

}
