package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// RoleHandler middleware checks if the user has the required role
func RoleHandler(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the user's role from the context (set during authentication)
		claims := c.MustGet("user").(jwt.MapClaims)
		admin, ok := claims["admin"].(bool)
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"status_code": http.StatusForbidden,
				"message":     "Forbidden: admin claim not found or invalid",
			})
			return
		}

		if (requiredRole == "admin" && !admin) || (requiredRole == "user" && admin) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"status_code": http.StatusForbidden,
				"message":     "Forbidden: insufficient permission",
			})
		}

		c.Next()
	}
}
