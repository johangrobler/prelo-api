package auth

import (
	"prelo/config"
	"time"

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

func GenerateToken(userID uint) (string, error) {
	/*
		claims2 := &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // Token valid for 24 hours
			Subject:   string(userID),
		}

	*/

	// Define the claims
	claims := CustomClaims{
		User: struct {
			ID        string `json:"id"`
			FirstName string `json:"firstName"`
			LastName  string `json:"lastName"`
			Role      string `json:"role"`
		}{
			ID:        "26bd8aab-979c-48f1-bfa1-13369fa7d061",
			FirstName: "Johan",
			LastName:  "Grobler",
			Role:      "TENANT",
		},
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	// Create the token using the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	jwtSecret := config.Config("JWT_SECRET")
	return token.SignedString([]byte(jwtSecret))
}

func ParseToken(tokenString string) (*jwt.RegisteredClaims, error) {

	jwtSecret := config.Config("JWT_SECRET")
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
