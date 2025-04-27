package main

import (
	"aesth-api/auth"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using passed environment")
	}

	router := gin.Default()
	router.POST("/auth", auth.AuthHandler)

	api := router.Group("/api")
	api.Use(auth.AuthMiddleware())
	{
		api.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"pong": "yeah"}) })
	}

	router.Run("0.0.0.0:8080")
}
