package dto

import (
	"net/http"

	orderDto "github.com/d02ev/ecommerce-api/internal/order/adapters/dto"
	sharedDto "github.com/d02ev/ecommerce-api/pkg/shared/dto"
)

type RegisterUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type RegisterUserResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginUserResponse struct {
	StatusCode   int    `json:"status_code"`
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

type RegisterUserDto struct {
	Name     string
	Email    string
	Password string
}

type LoginUserDto struct {
	Email    string
	Password string
}

type UserDto struct {
	Name      string                 `json:"name"`
	Email     string                 `json:"email"`
	Role      string                 `json:"role"`
	Addresses []sharedDto.AddressDto `json:"addresses"`
	Orders    []orderDto.OrderDto    `json:"orders"`
}

func NewRegisterUserResponse() *RegisterUserResponse {
	return &RegisterUserResponse{
		StatusCode: http.StatusCreated,
		Message:    "User registered successfully",
	}
}

func NewLoginUserResponse(accessToken, refreshToken string) *LoginUserResponse {
	return &LoginUserResponse{
		StatusCode:   http.StatusOK,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}
