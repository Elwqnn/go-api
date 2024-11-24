package controllers

import (
	"net/http"
	"user-service/services"

	"github.com/gin-gonic/gin"
)

// AuthController handles authentication requests
type AuthController struct {
	authService services.AuthService
}

// NewAuthController creates a new AuthController
func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

// Login handles user login and token generation
func (c *AuthController) Login(ctx *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	token, err := c.authService.Login(input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
