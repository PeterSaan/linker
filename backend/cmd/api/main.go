package main

import (
	"linker/app/controllers/auth"
	"linker/app/middleware"
	"linker/internal/database"
	"linker/internal/seeder"
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

    if err := seeder.Main(); err != nil {
        log.Fatalf("Error seeding database: %v", err)
    }  

    authRoutes := router.Group("/auth")
    {
        authRoutes.GET("/linkedin/callback", auth.OauthCallBack)
        authRoutes.GET("/linkedin/login", auth.OauthLogin)
        authRoutes.POST("/login", auth.PasswordLogin)
        authRoutes.POST("/logout", auth.Logout)
        authRoutes.POST("/register", auth.PasswordRegister)
    }

    authorized := router.Group("/", middleware.JWTMiddleware())
    {
        authorized.GET("/test")
    }


	router.Run(":" + port)
}
