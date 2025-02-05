package main

import (
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

	router.GET("/api/health", func(ctx *gin.Context) {
		chat := db.First(&models.Chat{})
		ctx.JSON(http.StatusOK, gin.H{
			"test": "working",
			"Chat": chat,
		})
	})

	router.Run(":" + port)
}
