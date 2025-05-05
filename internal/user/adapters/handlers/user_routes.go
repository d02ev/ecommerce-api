package handlers

import "github.com/gin-gonic/gin"

func RegisterAuthRoutes(rg *gin.RouterGroup, handler *AuthHandler) {
	auth := rg.Group("/auth");
	auth.POST("/register", handler.Register);
	auth.POST("/login", handler.Login);
}