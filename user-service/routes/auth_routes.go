package routes

import (
	"user-service/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterAuthRoutes registers authentication routes
func RegisterAuthRoutes(router *gin.Engine, authController *controllers.AuthController) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", authController.Login)
	}
}
