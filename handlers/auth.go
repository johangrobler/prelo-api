package handlers

import (
	"fmt"
	"prelo/config"
	"prelo/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func RegisterUser(c *fiber.Ctx) error {
	return models.RegisterUser(c)
}

// CustomClaims defines the structure of the JWT payload
type CustomClaims struct {
	User struct {
		ID        string `json:"id"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Role      string `json:"role"`
	} `json:"user"`
	jwt.RegisteredClaims
}

// Login route
func Login(c *fiber.Ctx) error {
	// Extract the credentials from the request body s
	loginRequest := new(models.LoginRequest)
	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// Find the user by credentials
	user, err := models.FindByCredentials(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// Define the claims

	claims := CustomClaims{
		User: struct {
			ID        string `json:"id"`
			FirstName string `json:"firstName"`
			LastName  string `json:"lastName"`
			Role      string `json:"role"`
		}{
			ID:        user.ID.String(),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Role:      "TENANT",
		},
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 240)),
		},
	}

	secret := config.Config("JWT_SECRET")
	fmt.Println(secret)
	//expirationTime := time.Now().Add(24 * time.Hour)

	// Create a new token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return err
	}
	// Return the token
	return c.Status(201).JSON(models.LoginResponse{
		AccessToken:  tokenString,
		RefreshToken: tokenString,
		UserID:       user.ID.String(),
	})
}

func Protected(c *fiber.Ctx) error {

	userID := c.Locals("userID").(string)
	return c.SendString("Welcome	userID:" + userID)

}
