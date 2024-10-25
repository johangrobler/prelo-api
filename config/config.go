package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

//const JWTSecret = "123qwe" // Replace with a strong secret key
//JWT_SECRET_KEY="123qwe"

// Config func to get env value from key ---
func Config(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}
