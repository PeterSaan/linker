package middleware

import (
	"linker/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        tokenString, err := ctx.Cookie("linker_token")
        if err != nil {
            ctx.JSON(http.StatusUnauthorized, gin.H{
                "error": "No token found",
            })
            ctx.Abort()
            return
        }

        token, err := services.ParseToken(tokenString)
        if err != nil {
            ctx.JSON(http.StatusUnauthorized, gin.H{
                "error": err.Error(),
            })
            ctx.Abort()
            return
        }

        ctx.Set("user_id", token.UserId)

        ctx.Next()
    }
}
