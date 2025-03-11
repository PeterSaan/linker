package auth

import (
	"linker/app/services"
	"linker/internal/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PasswordRegister(ctx *gin.Context) {
    var formData structs.PasswordRegiserData

    if err := ctx.ShouldBindJSON(&formData); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    if formData.Password != formData.PasswordConfirm {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "Both password fields must match",
        })
        return
    }

    user, err := services.PasswordRegisterUser(formData) 
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    tokenString, err := services.GenerateToken(&user)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    services.SetTokenCookie(ctx, tokenString)
}

func PasswordLogin(ctx *gin.Context) {
    var formData structs.PasswordLoginData

    if err := ctx.ShouldBindJSON(&formData); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    user, err := services.PasswordGetUser(formData)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    tokenString, err := services.GenerateToken(&user)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    services.SetTokenCookie(ctx, tokenString)
}

func Logout(ctx *gin.Context) {
    ctx.SetCookie("linker_token", "", -1, "/", "localhost", false, true)
    ctx.JSON(http.StatusOK, gin.H{
        "success": true,
        "message": "Logged out successfully!",
    })
}
