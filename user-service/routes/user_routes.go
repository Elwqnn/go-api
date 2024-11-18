package routes

import (
	"user-service/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes registers user-related routes
func RegisterUserRoutes(router *gin.Engine, userController *controllers.UserController) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("/", userController.CreateUser)
		userGroup.GET("/", userController.GetAllUsers)
		userGroup.GET("/:id", userController.GetUserByID)
		userGroup.PUT("/:id", userController.UpdateUser)
		userGroup.DELETE("/:id", userController.DeleteUser)
	}
}
