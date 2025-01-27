package domain

import "github.com/golang-jwt/jwt/v5"

type JWTClaims struct {
	UserID   string `json:"sub"`
	Username string `json:"name"`
	jwt.RegisteredClaims
}
