package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	"log"

	"go-api/config"
	"go-api/internal/auth"
	"go-api/internal/controllers"
	"go-api/internal/models"
	"go-api/internal/repositories"
	"go-api/internal/routes"
	"go-api/internal/services"
)

// @title Go API
// @version 1.0
// @description This is a sample server for a Go API.
// @host localhost:8080
// @BasePath /

func main() {
	config.LoadConfig()
	router := setupRouter()
	port := config.AppConfig.Server.Port
	err := router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	db := setupDatabase()
	userController := initializeUserController(db)

	// Public route for login
	router.POST("/auth/login", userController.Login)

	// Protected routes
	protected := router.Group("/")
	protected.Use(auth.AuthMiddleware())
	routes.RegisterUserRoutes(protected, userController)

	// Swagger endpoint (e.g: http://localhost:8080/swagger/index.html)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}

func setupDatabase() *gorm.DB {
	db := config.ConnectDB()
	models.MigrateDB(db)
	return db
}

func initializeUserController(db *gorm.DB) *controllers.UserController {
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	return controllers.NewUserController(userService)
}
