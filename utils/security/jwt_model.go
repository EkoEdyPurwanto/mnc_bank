package security

import "github.com/golang-jwt/jwt/v5"

type AppClaims struct {
	jwt.RegisteredClaims
	Name string `json:"name"`
}
