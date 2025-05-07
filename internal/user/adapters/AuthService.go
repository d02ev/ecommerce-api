package adapters

import (
	"errors"
	"strings"

	"github.com/d02ev/ecommerce-api/internal/user/adapters/dto"
	"github.com/d02ev/ecommerce-api/internal/user/domain"
	"github.com/d02ev/ecommerce-api/internal/user/ports"
	"github.com/d02ev/ecommerce-api/pkg/custom_errors"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type AuthService struct {
	userRepository  ports.IUserRepository
	passwordService ports.IPasswordService
	tokenService    ports.ITokenService
}

func NewAuthService(userRepository ports.IUserRepository, passwordService ports.IPasswordService, tokenService ports.ITokenService) *AuthService {
	return &AuthService{
		userRepository:  userRepository,
		passwordService: passwordService,
		tokenService:    tokenService,
	}
}

func (as *AuthService) RegisterUser(registerUserDto dto.RegisterUserDto) (*dto.RegisterUserResponse, error) {
	existingUser, err := as.userRepository.FindByEmail(registerUserDto.Email)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, custom_errors.ErrInternalServerError
		}
	}

	if existingUser != nil {
		return nil, domain.ErrUserAlreadyExists
	}

	hashedPassword, err := as.passwordService.Hash(registerUserDto.Password)
	if err != nil {
		return nil, custom_errors.ErrInternalServerError
	}

	var role uint = 0
	if strings.Contains(registerUserDto.Email, "admin") {
		role = 1
	}

	user := &domain.UserEntity{
		Name:         registerUserDto.Name,
		Email:        registerUserDto.Email,
		PasswordHash: hashedPassword,
		Role:         role,
	}

	err = as.userRepository.Save(user)
	if err != nil {
		return nil, custom_errors.ErrInternalServerError
	}

	return dto.NewRegisterUserResponse(), nil
}

func (as *AuthService) LoginUser(loginUserDto dto.LoginUserDto) (*dto.LoginUserResponse, error) {
	user, err := as.userRepository.FindByEmail(loginUserDto.Email)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, custom_errors.ErrInternalServerError
		}
		return nil, domain.ErrUserNotFound
	}

	if !as.passwordService.Compare(user.PasswordHash, loginUserDto.Password) {
		return nil, domain.ErrInvalidCredentials
	}

	accessToken, err := as.tokenService.GenerateAccessToken(user.ID, user.Role)
	if err != nil {
		return nil, custom_errors.ErrInternalServerError
	}
	refreshToken, err := as.tokenService.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, custom_errors.ErrInternalServerError
	}

	err = as.userRepository.UpdateRefreshToken(user.ID, refreshToken)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, custom_errors.ErrInternalServerError
		}
		return nil, domain.ErrUserNotFound
	}

	return dto.NewLoginUserResponse(accessToken, refreshToken), nil
}

func (as *AuthService) RefreshAccessToken(refreshToken string) (string, error) {
	userId, err := as.tokenService.DecodeRefreshToken(refreshToken)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return "", domain.ErrTokenExpired
		} else {
			return "", domain.ErrValidatingToken
		}
	}

	user, err := as.userRepository.FindByID(userId)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return "", custom_errors.ErrInternalServerError
		}
		return "", domain.ErrUserNotFound
	}

	accessToken, err := as.tokenService.GenerateAccessToken(user.ID, user.Role)
	if err != nil {
		return "", custom_errors.ErrInternalServerError
	}

	return accessToken, nil
}
