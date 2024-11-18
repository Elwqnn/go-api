package main

import (
	"log"
	"user-service/controllers"
	"user-service/repositories"
	"user-service/routes"
	"user-service/services"

	"github.com/Elwqnn/go-api/shared-utils/database"
	"github.com/Elwqnn/go-api/shared-utils/env"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	env.LoadEnv()

	// Initialize database
	dsn := "postgres://username:password@localhost:5432/userdb?sslmode=disable"
	database.ConnectDB(dsn)

	// Dependency injection
	db := database.GetDB()
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// Setup Gin router
	router := gin.Default()
	routes.RegisterUserRoutes(router, userController)

	// Start the server
	log.Println("User service running on port 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
