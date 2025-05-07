package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrTokenMalformed
		}
		return []byte(viper.GetString("ACCESS_TOKEN_SECRET")), nil
	})
}

func AuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		parts := strings.SplitN(authHeader, " ", 2)

		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status_code": http.StatusUnauthorized,
				"message":     "missing or malformed auth token",
			})
			return
		}

		tokenString := parts[1]
		fmt.Print("Token: ", tokenString, "\n")
		token, err := ValidateToken(tokenString)
		if err != nil || !token.Valid {
			if errors.Is(err, jwt.ErrTokenExpired) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"status_code": http.StatusUnauthorized,
					"message":     "token expired",
				})
			} else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"status_code": http.StatusUnauthorized,
					"message":     "invalid token",
				})
			}

			return
		}

		fmt.Print("Token claims: ", token.Claims.(jwt.MapClaims), "\n")

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("user", claims)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status_code": http.StatusUnauthorized,
				"message":     "cannot parse claims",
			})
			return
		}
	}
}
