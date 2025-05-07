package ports

import "github.com/d02ev/ecommerce-api/internal/user/domain"

type IUserRepository interface {
	Save(user *domain.UserEntity) error
	FindByID(id uint) (*domain.UserEntity, error)
	FindByEmail(email string) (*domain.UserEntity, error)
	UpdateRefreshToken(userId uint, refreshToken string) error
}
