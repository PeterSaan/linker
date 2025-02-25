package services

import (
	"fmt"
	"linker/internal/models"
	"linker/internal/structs"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var key = []byte(os.Getenv("KEY"))

func GenerateToken(user *models.User) (string, error) {
	expirationTime := time.Now().AddDate(0, 0, 30)

	claims := &structs.JWTClaims{
		UserId: int(user.ID),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		fmt.Println("Error creating JWT token: ", err)
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*structs.JWTClaims, error) {
	claims := &structs.JWTClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("Unautorized")
	}

	return claims, err
}

func SetTokenCookie(ctx *gin.Context, tokenString string) {
    expiration := 60 * 60 * 24 * 30

    ctx.SetCookie("linker_token", tokenString, expiration, "/", "localhost", false, true)
}
