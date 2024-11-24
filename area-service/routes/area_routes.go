package routes

import (
	"area-service/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterAreaRoutes registers area-related routes
func RegisterAreaRoutes(protected *gin.RouterGroup, areaController *controllers.AreaController) {
	userGroup := protected.Group("/areas")
	{
		userGroup.POST("/", areaController.CreateArea)
		userGroup.GET("/", areaController.GetAllAreas)
		userGroup.GET("/:id", areaController.GetAreaByID)
		userGroup.PUT("/:id", areaController.UpdateArea)
		userGroup.DELETE("/:id", areaController.DeleteArea)
	}
}
