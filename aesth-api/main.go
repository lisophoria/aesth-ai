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
	api := router.Group("/api")
	{
		api.POST("/auth", auth.AuthHandler)

		protected := router.Group("")
		protected.Use(auth.AuthMiddleware())
		protected.GET("/ping", func(c *gin.Context) { c.IndentedJSON(http.StatusOK, "pong") })
	}

	router.Run("localhost:8080")
}
