package ports

import "github.com/d02ev/ecommerce-api/internal/product/domain"

type IProductRepository interface {
	Save(productEntity *domain.ProductEntity) error
	FindByID(id uint) (*domain.ProductEntity, error)
	GetAll() ([]*domain.ProductEntity, error)
	Update(id uint, productEntity *domain.ProductEntity) error
	Delete(id uint) error
}
