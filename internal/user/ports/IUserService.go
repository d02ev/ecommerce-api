package ports

import (
	"github.com/d02ev/ecommerce-api/internal/user/adapters/dto"
)

type IUserService interface {
	GetUserDetails(userId uint) (*dto.UserDto, error)
}
