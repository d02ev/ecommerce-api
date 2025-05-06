package adapters

import (
	"errors"

	"github.com/d02ev/ecommerce-api/internal/user/domain"
	"github.com/d02ev/ecommerce-api/internal/user/ports"
	"gorm.io/gorm"
)

type UserService struct {
	UserRepository ports.IUserRepository
}

func NewUserService(userRepository ports.IUserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (us *UserService) GetUserDetails(userId uint) (*domain.UserEntity, error) {
	user, err := us.UserRepository.FindByID(userId); if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	userEntity := &domain.UserEntity{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Role: user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Addresses: user.Addresses,

	}

	return userEntity, nil
}