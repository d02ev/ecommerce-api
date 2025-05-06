package handlers

import (
	"errors"
	"net/http"

	"github.com/d02ev/ecommerce-api/internal/user/adapters/dto"
	"github.com/d02ev/ecommerce-api/internal/user/domain"
	"github.com/d02ev/ecommerce-api/internal/user/ports"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type CookieOptions struct {
	Path 	 string
	Domain 	 string
	Secure 	 bool
	HttpOnly bool
}

type AuthHandler struct {
	authService ports.IAuthService
	CookieOptions
}

func NewAuthHandler(authService ports.IAuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		CookieOptions: CookieOptions{
			Path: 	 viper.GetString("COOKIE_PATH"),
			Domain: 	 viper.GetString("COOKIE_DOMAIN"),
			Secure: 	 viper.GetBool("COOKIE_SECURE"),
			HttpOnly: viper.GetBool("COOKIE_HTTP_ONLY"),
		},
	}
}

func (ah *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterUserRequest;
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"message":     "Invalid request",
			"error": 		 err.Error(),
		})
	}

	res, err := ah.authService.RegisterUser(req.Name, req.Email, req.Password);
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

func (ah *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginUserRequest;
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"message":     "Invalid request",
			"error": 		 err.Error(),
		})
		return
	}

	res, err := ah.authService.LoginUser(req.Email, req.Password);
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

	c.SetCookie("refresh_token", res.RefreshToken, 7 * 24 * 60 * 60, ah.Path, ah.Domain, ah.Secure, ah.HttpOnly);
	c.SetCookie("access_token", res.AccessToken, 60 * 60 * 60, ah.Path, ah.Domain, ah.Secure, ah.HttpOnly);
	c.JSON(http.StatusOK, res);
}

func (ah *AuthHandler) RefreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token");
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status_code": http.StatusUnauthorized,
			"message":     "Refresh token not found",
		})
		return
	}

	res, err := ah.authService.RefreshAccessToken(refreshToken);
	if err != nil {
		if errors.Is(err, domain.ErrTokenExpired) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status_code": http.StatusUnauthorized,
				"message":     "Refresh token expired",
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

	c.SetCookie("access_token", res, 60 * 60 * 60, ah.Path, ah.Domain, ah.Secure, ah.HttpOnly);
	c.SetCookie("refresh_token", refreshToken, 7 * 24 * 60 * 60, ah.Path, ah.Domain, ah.Secure, ah.HttpOnly);
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message": 	 "Access token refreshed",
	});
}

func (ah *AuthHandler) Logout(c *gin.Context) {
	c.SetCookie("refresh_token", "", -1, ah.Path, ah.Domain, ah.Secure, ah.HttpOnly);
	c.SetCookie("access_token", "", -1, ah.Path, ah.Domain, ah.Secure, ah.HttpOnly);
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message": 	 "Logged out successfully",
	});
}