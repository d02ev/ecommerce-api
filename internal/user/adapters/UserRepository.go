package adapters

import (
	"github.com/d02ev/ecommerce-api/internal/user/domain"
	"github.com/d02ev/ecommerce-api/pkg/models"
	"github.com/d02ev/ecommerce-api/pkg/shared/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Save(userEntity *domain.UserEntity) error {
	// convert the domain entity to the gorm model
	user := &models.User{
		Name:         userEntity.Name,
		Email:        userEntity.Email,
		Role:         userEntity.Role,
		PasswordHash: userEntity.PasswordHash,
		RefreshToken: userEntity.RefreshToken,
	}

	// try to save the user
	res := ur.db.Save(user)
	if res.Error != nil {
		return res.Error
	}

	// if saved successfully return nil
	return nil
}

func (ur *UserRepository) FindByID(id uint) (*domain.UserEntity, error) {
	var user models.User

	if err := ur.db.Preload("Addresses").First(&user, id).Error; err != nil {
		return nil, err
	}

	userEntity := &domain.UserEntity{
		ID:           user.ID,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		Name:         user.Name,
		Email:        user.Email,
		Role:         user.Role,
		PasswordHash: user.PasswordHash,
		RefreshToken: user.RefreshToken,
		Addresses:    make([]entities.AddressEntity, len(user.Addresses)),
	}

	return userEntity, nil
}

func (ur *UserRepository) FindByEmail(email string) (*domain.UserEntity, error) {
	var user models.User

	if err := ur.db.Preload("Addresses").Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	userEntity := &domain.UserEntity{
		ID:           user.ID,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		Name:         user.Name,
		Email:        user.Email,
		Role:         user.Role,
		PasswordHash: user.PasswordHash,
		RefreshToken: user.RefreshToken,
		Addresses:    make([]entities.AddressEntity, len(user.Addresses)),
	}

	return userEntity, nil
}

func (ur *UserRepository) UpdateRefreshToken(userId uint, refreshToken string) error {
	var user models.User

	if err := ur.db.Model(&user).Where("id = ?", userId).Update("refresh_token", refreshToken).Error; err != nil {
		return err
	}

	return nil
}
