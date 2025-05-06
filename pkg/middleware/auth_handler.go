package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		accessTokenSecretKey := viper.GetString("ACCESS_TOKEN_SECRET");
		return []byte(accessTokenSecretKey), nil
	})
}

func AuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization");

		if authHeader == "" {
			c.JSON(401, gin.H{
				"status_code": 401,
				"message":     "Unauthorized",
			})
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ");
		_, err := ValidateToken(token); if err != nil {
			c.JSON(401, gin.H{
				"status_code": 401,
				"message":     "Unauthorized",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}