package main

import (
	"linker/app/controllers/auth"
	"linker/internal/database"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	router := gin.Default()
	port := os.Getenv("GIN_PORT")
	if port == "" {
		port = "8080"
	}

	_, err := database.Connect()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

    authentication := router.Group("/auth")
    {
        authentication.GET("/linkedin/callback", auth.OauthCallBack)
        authentication.GET("/linkedin/login", auth.OauthLogin)
        authentication.POST("/login", auth.PasswordLogin)
        authentication.POST("/logout", auth.Logout)
        authentication.POST("/register", auth.PasswordRegister)
    }


	router.Run(":" + port)
}
