package middlewares

import (
	"fmt"
	//"prelo/auth"
	"prelo/auth"
	"prelo/config"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

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

// Middleware JWT function
func NewAuthMiddleware() fiber.Handler {

	jwtSecret := config.Config("JWT_SECRET")
	fmt.Println(jwtSecret)

	return jwtware.New(jwtware.Config{
		SigningKey:    []byte(jwtSecret), //[]byte("fRQg7MzjNFEjqUZ2nG5azPMZRnM5Gtn7uh8QQMymF56e7jAHkdBfgfrfu7ETwvzpQXMkuRa3PHXtZfQV4sp79epK94RmcZ4Z83WLhergnfw85MsdT48MhVt6J8dB3MEg32q8wVXF7fA2ymsunAYDSJfXwkgV8kfyQt3wtUpvqQp3WHhzNGG4KjQP2wFUzXYkF9BarZFuCCVxtfBGYh4v6A4m3EVhXz7M5s7be3UHek9hybxTuvDF9tsFgCsfp2Kn"),
		SigningMethod: "HS256",
		Claims:        &auth.CustomClaims{},

		SuccessHandler: func(c *fiber.Ctx) error {

			// Extract the claims from the token
			user := c.Locals("user").(*jwt.Token)
			claims := user.Claims.(*auth.CustomClaims)

			// You can now use the claims in your route handler
			fmt.Println("User ID:", claims.User.ID)
			c.Locals("userID", claims.User.ID)
			c.Locals("userRole", claims.User.Role)

			return c.Next()
		},

		ErrorHandler: func(c *fiber.Ctx, err error) error {

			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "Unauthorized",
				})
			}

			return c.Next()
		}})

}
