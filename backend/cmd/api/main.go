package main

import (
	"linker/internal/auth"
	"linker/internal/database"
	"linker/internal/models"
	"log"
	"net/http"
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

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	result := map[string]interface{}{}
	log.Printf("User found: %+v\n", db.Model(&models.Profile{}).First(&result))

	router.GET("/api/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"test": "working",
			"pp":   result,
		})
	})

    authentication := router.Group("/auth")
    {
        authentication.GET("/linkedin/callback", auth.OauthCallBack)

        authentication.GET("/linkedin/login", auth.OauthLogin)

        authentication.GET("/login") // login without linkedin
    }


	router.Run(":" + port)
}
