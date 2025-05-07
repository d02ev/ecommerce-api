package ports

import (
	"github.com/d02ev/ecommerce-api/internal/user/adapters/dto"
)

type IAuthService interface {
	RegisterUser(registerUserDto dto.RegisterUserDto) (*dto.RegisterUserResponse, error)
	LoginUser(loginUserDto dto.LoginUserDto) (*dto.LoginUserResponse, error)
	RefreshAccessToken(refreshToken string) (string, error)
}
