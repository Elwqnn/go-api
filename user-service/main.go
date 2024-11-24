package main

import (
	"log"
	"user-service/controllers"
	"user-service/middlewares"
	"user-service/repositories"
	"user-service/routes"
	"user-service/services"

	"github.com/gin-gonic/gin"
	"user-service/shared-utils/database"
	"user-service/shared-utils/env"
)

func main() {
	env.LoadEnv()

	database.ConnectDB()

	db := database.GetDB()
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	authService := services.NewAuthService(userRepo)
	userController := controllers.NewUserController(userService)
	authController := controllers.NewAuthController(authService)

	// Setup router
	r := gin.Default()

	// Public routes
	routes.RegisterAuthRoutes(r, authController)

	// Protected routes
	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddleware())
	routes.RegisterUserRoutes(protected, userController)

	// Start the server
	log.Println("User service running on port 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
