package ports

import "github.com/d02ev/ecommerce-api/internal/user/domain"

type IUserService interface {
	GetUserDetails(userId uint) (*domain.UserEntity, error)
}