package structs

import "github.com/golang-jwt/jwt/v5"

type JWTClaims struct {
	UserId int `json:"user_id"`
    jwt.RegisteredClaims
}
