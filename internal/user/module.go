package user

import (

	"github.com/d02ev/ecommerce-api/internal/user/adapters"
	"github.com/d02ev/ecommerce-api/internal/user/adapters/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitUserModule(router *gin.RouterGroup, db *gorm.DB) {
	userRepository := adapters.NewUserRepository(db);
	passwordService := adapters.NewPasswordService();
	tokenService := adapters.NewTokenService();
	authService := adapters.NewAuthService(userRepository, passwordService, tokenService);
	userService := adapters.NewUserService(userRepository);
	authHandler := handlers.NewAuthHandler(authService);
	userHandler := handlers.NewUserHandler(userService);
	handlers.RegisterAuthRoutes(router, authHandler);
	handlers.RegisterUserRoutes(router, userHandler);
}