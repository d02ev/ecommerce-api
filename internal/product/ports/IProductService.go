package ports

import (
	"github.com/d02ev/ecommerce-api/internal/product/adapters/dto"
)

type IProductService interface {
	Create(createProductDto dto.CreateOrUpdateProductDto) (*dto.CreateProductResponse, error)
	FetchById(id uint) (*dto.FetchProductResponse, error)
	FetchAll() ([]*dto.FetchProductResponse, error)
	Update(updateProductDto dto.CreateOrUpdateProductDto) (*dto.UpdateProductResponse, error)
	Delete(id uint) (*dto.DeleteProductResponse, error)
}
