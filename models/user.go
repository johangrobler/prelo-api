package models

import (
	"errors"
	"prelo/database"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	UserID       string `json:"userId"`
}

// User struct
type User struct {
	gorm.Model
	ID           uuid.UUID `gorm:"type:uuid;"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email" gorm:"uniqueIndex"`
	Mobile       string    `json:"mobile"`
	Password     string    `json:"password"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
}

// Users struct
type Users struct {
	Users []User `json:"users"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	user.ID = uuid.New()
	return
}

func GetUserProfile(c *fiber.Ctx) error {

	userID := c.Locals("userID").(string)

	db := database.DB.Db
	var model User

	db.Where("id = ?", userID).First(&model)

	// Return the created user
	return c.Status(201).JSON(fiber.Map{"status": "success", "user": model})

}

// Simulate a database call
func FindByCredentials(email, password string) (*User, error) {
	if email == "" {

		return nil, errors.New("credentials not found")
	}

	db := database.DB.Db
	// get id params
	var user User

	result := db.Find(&user, "email = ?", email)

	// Check for errors
	if result.Error != nil {
		return nil, result.Error
	}
	// Here you would query your database for the user with the given email
	if email == user.Email && password == user.Password {
		return &user, nil
	}
	return nil, errors.New("user not found")
}

// Register user
func RegisterUser(c *fiber.Ctx) error {
	db := database.DB.Db
	m := new(User)
	// Store the body in the user and return error if encountered
	err := c.BodyParser(m)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	var user User

	result := db.Find(&user, "email = ?", m.Email)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	}
	// Here you would query your database for the user with the given email
	if m.Email == user.Email {
		return c.Status(409).JSON(fiber.Map{"status": "error", "message": "Email has been taken"})

	}

	err = db.Create(&m).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	}
	// Return the created user
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "User has created", "data": m})
}
