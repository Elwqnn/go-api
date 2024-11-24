package middlewares

import (
	"net/http"
	"user-service/utils"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware protects routes that require authentication
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		token := authHeader[len("Bearer "):]
		claims, err := utils.ValidateJWT(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Set user ID in context for use in handlers
		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
