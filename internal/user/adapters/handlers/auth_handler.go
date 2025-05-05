package handlers

import (
	"errors"
	"net/http"

	"github.com/d02ev/ecommerce-api/internal/user/adapters/dto"
	"github.com/d02ev/ecommerce-api/internal/user/domain"
	"github.com/d02ev/ecommerce-api/internal/user/ports"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService ports.IAuthService
}

func NewAuthHandler(authService ports.IAuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (uh *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterUserRequest;
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"message":     "Invalid request",
			"error": 		 err.Error(),
		})
	}

	res, err := uh.authService.RegisterUser(req.Name, req.Email, req.Password);
	if err != nil {
		if errors.Is(err, domain.ErrUserAlreadyExists) {
			c.JSON(http.StatusConflict, gin.H{
				"status_code": http.StatusConflict,
				"message":     "User already exists",
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status_code": http.StatusInternalServerError,
				"message":     "Internal server error",
			})
			return
		}
	}

	c.JSON(http.StatusCreated, res);
}

func (uh *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginUserRequest;
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"message":     "Invalid request",
			"error": 		 err.Error(),
		})
		return
	}

	res, err := uh.authService.LoginUser(req.Email, req.Password);
	if err != nil {
		if errors.Is(err, domain.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status_code": http.StatusUnauthorized,
				"message":     "Invalid credentials",
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status_code": http.StatusInternalServerError,
				"message":     "Internal server error",
			})
			return
		}
	}

	c.JSON(http.StatusOK, res);
}