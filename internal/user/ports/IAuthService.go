package ports

import (
	"github.com/d02ev/ecommerce-api/internal/user/adapters/dto"
)

type IAuthService interface {
	RegisterUser(name, email, password string) (*dto.RegisterUserResponse, error)
	LoginUser(email, password string) (*dto.LoginUserResponse, error)
	RefreshAccessToken(refreshToken string) (string, error)
}
