package handlers

import (
	"net/http"
	"strconv"

	"github.com/d02ev/ecommerce-api/internal/user/ports"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserHandler struct {
	userService ports.IUserService
}

func NewUserHandler(userService ports.IUserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (uh *UserHandler) Me(c *gin.Context) {
	claims := c.MustGet("user").(jwt.MapClaims)
	subject, err := claims.GetSubject()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status_code": http.StatusUnauthorized,
			"message":     "Unauthorized",
		})
	}

	userId, err := strconv.Atoi(subject)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"message":     "Invalid user ID",
		})
	}

	user, err := uh.userService.GetUserDetails(uint(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"message":     "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     "User details retrieved successfully",
		"user":        user,
	})
}
