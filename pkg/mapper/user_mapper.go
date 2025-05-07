package mapper

import (
	"github.com/d02ev/ecommerce-api/internal/user/adapters/dto"
	"github.com/d02ev/ecommerce-api/internal/user/domain"
	sharedDto "github.com/d02ev/ecommerce-api/pkg/shared/dto"
	"github.com/d02ev/ecommerce-api/pkg/shared/entities"
)

func mapAddressEntityToDto(addressEntity entities.AddressEntity) sharedDto.AddressDto {
	return sharedDto.AddressDto{
		AddLine1: addressEntity.AddLine1,
		AddLine2: addressEntity.AddLine2,
		Landmark: addressEntity.Landmark,
		City:     addressEntity.City,
		Country:  addressEntity.Country,
		ZipCode:  addressEntity.ZipCode,
	}
}

func mapAddressEntitiesToDtos(addressEntities []entities.AddressEntity) []sharedDto.AddressDto {
	addressDtos := make([]sharedDto.AddressDto, 0, len(addressEntities))

	for _, addressEntity := range addressEntities {
		addressDto := mapAddressEntityToDto(addressEntity)
		addressDtos = append(addressDtos, addressDto)
	}

	return addressDtos
}

func MapUserEntityToDto(userEntity *domain.UserEntity) *dto.UserDto {
	var role string = "user"
	if userEntity.Role == 1 {
		role = "admin"
	}
	return &dto.UserDto{
		Name:      userEntity.Name,
		Email:     userEntity.Email,
		Role:      role,
		Addresses: mapAddressEntitiesToDtos(userEntity.Addresses),
	}
}
