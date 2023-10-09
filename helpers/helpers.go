package helpers

import (
	"log"
	models "mini-wallet/models"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func GenerateToken(customer_xid string, isEnabled bool) (string, error) {
	claims := &models.Claims{
		Customer_XID: customer_xid,
		IsEnabled:    isEnabled,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("SECRET"))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func GoDotENVLoader(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(key)
}
