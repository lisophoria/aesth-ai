package main

import (
	"aesth-api/auth"
	"aesth-api/database"
	"aesth-api/handlers"
	"aesth-api/repositories"
	"aesth-api/services"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using passed environment")
	}

	db := database.Init()

	userRepository := repositories.NewUserRepository(db)
	relevancyRepository := repositories.NewRelevancyRepository(db)

	jwtService := auth.NewJwtService()
	relevancyService := services.NewRelevancyService(relevancyRepository)
	authService := services.NewAuthService(userRepository, relevancyRepository)

	userHandler := handlers.NewUserHandler(userRepository)
	authHandler := handlers.NewAuthHandler(jwtService, authService)
	relevancyHandler := *handlers.NewRelevancyHandler(relevancyService)

	router := gin.Default()
	authRouter := router.Group("/auth")
	{
		authRouter.POST("/login", authHandler.Login)
		authRouter.POST("/register", authHandler.Register)
	}

	apiRouter := router.Group("/api")
	apiRouter.Use(auth.AuthMiddleware(jwtService))
	{
		usersRouter := apiRouter.Group("/users")
		{
			usersRouter.GET("/", userHandler.GetUsers)
			usersRouter.GET("/:id", userHandler.GetUser)
			usersRouter.PUT("/:id", userHandler.UpdateUser)
			usersRouter.DELETE(":id", userHandler.DeleteUser)
		}

		relevancyRouter := apiRouter.Group("/relevancy")
		{
			relevancyRouter.GET("/:user_id", relevancyHandler.GetByUserID)
			relevancyRouter.POST("/:user_id/adjust", relevancyHandler.AdjustRelevancy)
			relevancyRouter.POST("/:user_id/get-pair", relevancyHandler.GetRelevancyPair)
		}
	}

	router.Run("0.0.0.0:8080")
}
