package main

import (
	"area-service/controllers"
	"area-service/repositories"
	"area-service/routes"
	"area-service/services"
	"area-service/shared-utils/database"
	"area-service/shared-utils/env"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	env.LoadEnv()

	database.ConnectDB()

	db := database.GetDB()
	areaRepo := repositories.NewAreaRepository(db)
	areaService := services.NewAreaService(areaRepo)
	areaController := controllers.NewAreaController(areaService)

	// Setup router
	r := gin.Default()

	// Public routes
	routes.RegisterAreaRoutes(r, areaController)

	// Protected routes
	protected := r.Group("/")
	routes.RegisterAreaRoutes(protected, areaController)

	// Start the server
	log.Println("User service running on port 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
