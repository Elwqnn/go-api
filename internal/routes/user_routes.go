package routes

import (
	"github.com/gin-gonic/gin"

	"go-api/internal/controllers"
)

// RegisterUserRoutes sets up user-related routes
func RegisterUserRoutes(protected *gin.RouterGroup, userController *controllers.UserController) {
	protected.POST("/users", userController.CreateUser)
	protected.GET("/users", userController.GetUsers)
	protected.GET("/users/:id", userController.GetUserByID)
	protected.PUT("/users/:id", userController.UpdateUser)
	protected.DELETE("/users/:id", userController.DeleteUser)
}
