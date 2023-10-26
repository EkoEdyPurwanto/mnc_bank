package security

import (
	"EkoEdyPurwanto/mnc-bank/model"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// GenerateJWTToken generates a JWT token for the customer
func GenerateJWTToken(customer *model.Customer) (string, error) {
	// Create a new token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the claims (payload) for the token
	claims := token.Claims.(jwt.MapClaims)
	claims["ID"] = customer.ID
	claims["exp"] = time.Now().Add(3 * time.Hour)

	// Sign the token with a secret key
	secretKey := []byte("rhs") // Replace with your secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
