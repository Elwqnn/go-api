package controllers

import (
	"net/http"
	"strconv"

	"area-service/models"
	"area-service/services"
	"github.com/gin-gonic/gin"
)

type AreaController struct {
	service services.AreaService
}

func NewAreaController(service services.AreaService) *AreaController {
	return &AreaController{service: service}
}

func (c *AreaController) CreateArea(ctx *gin.Context) {
	var area models.Area
	if err := ctx.ShouldBindJSON(&area); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := c.service.CreateArea(&area); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, area)
}

func (c *AreaController) GetAreaByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	area, err := c.service.GetAreaByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Area not found"})
		return
	}

	ctx.JSON(http.StatusOK, area)
}
