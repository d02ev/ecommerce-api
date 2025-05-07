package adapters

import (
	"errors"

	"github.com/d02ev/ecommerce-api/internal/user/adapters/dto"
	"github.com/d02ev/ecommerce-api/internal/user/domain"
	"github.com/d02ev/ecommerce-api/internal/user/ports"
	"github.com/d02ev/ecommerce-api/pkg/mapper"
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

func (us *UserService) GetUserDetails(userId uint) (*dto.UserDto, error) {
	user, err := us.UserRepository.FindByID(userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	userDto := mapper.MapUserEntityToDto(user)

	return userDto, nil
}
