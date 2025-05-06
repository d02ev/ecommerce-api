package handlers

import (
	"github.com/d02ev/ecommerce-api/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(rg *gin.RouterGroup, handler *AuthHandler) {
	auth := rg.Group("/auth");
	auth.POST("/register", handler.Register);
	auth.POST("/login", handler.Login);
	auth.POST("/refresh-token", handler.RefreshToken);
	auth.POST("/logout", middleware.AuthHandler(), handler.Logout);
}

func RegisterUserRoutes(rg *gin.RouterGroup, handler *UserHandler) {
	user := rg.Group("/user");
	user.GET("/me", middleware.AuthHandler(), handler.Me);
}