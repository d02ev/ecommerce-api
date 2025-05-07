package adapters

import (
	"github.com/d02ev/ecommerce-api/internal/product/domain"
	"github.com/d02ev/ecommerce-api/pkg/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (pr *ProductRepository) Save(productEntity *domain.ProductEntity) error {
	// Check if the category exists
	var category models.Category
	if err := pr.db.First(&category, productEntity.CategoryID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Create a new category if it doesn't exist
			newCategory := &models.Category{
				Name:        productEntity.Category.Name,
				Description: productEntity.Category.Description,
			}
			if err := pr.db.Create(newCategory).Error; err != nil {
				return err
			}
			// Assign the new category's ID to the product
			productEntity.CategoryID = newCategory.ID
		} else {
			return err
		}
	}

	product := &models.Product{
		Name:        productEntity.Name,
		Description: productEntity.Description,
		Price:       productEntity.Price,
		SKU:         productEntity.SKU,
		StockQty:    productEntity.StockQty,
		CategoryID:  productEntity.CategoryID,
		Category:    category,
	}

	// Save the product
	if err := pr.db.Create(product).Error; err != nil {
		return err
	}

	return nil
}

func (pr *ProductRepository) FindByID(id uint) (*domain.ProductEntity, error) {
	var product models.Product

	if err := pr.db.Preload("Category").First(&product, id).Error; err != nil {
		return nil, err
	}

	categoryEntity := domain.CategoryEntity{
		ID:          product.CategoryID,
		CreatedAt:   product.Category.CreatedAt,
		UpdatedAt:   product.Category.UpdatedAt,
		Name:        product.Category.Name,
		Description: product.Category.Description,
	}
	productEntity := &domain.ProductEntity{
		ID:          product.ID,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		SKU:         product.SKU,
		StockQty:    product.StockQty,
		CategoryID:  product.CategoryID,
		Category:    categoryEntity,
	}

	return productEntity, nil
}

func (pr *ProductRepository) GetAll() ([]*domain.ProductEntity, error) {
	var products []models.Product

	// Fetch all products and preload the associated categories
	if err := pr.db.Preload("Category").Find(&products).Error; err != nil {
		return nil, err
	}

	var productEntities []*domain.ProductEntity
	for _, product := range products {
		categoryEntity := domain.CategoryEntity{
			ID:          product.Category.ID,
			CreatedAt:   product.Category.CreatedAt,
			UpdatedAt:   product.Category.UpdatedAt,
			Name:        product.Category.Name,
			Description: product.Category.Description,
		}

		productEntity := &domain.ProductEntity{
			ID:          product.ID,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			SKU:         product.SKU,
			StockQty:    product.StockQty,
			CategoryID:  product.CategoryID,
			Category:    categoryEntity,
		}

		productEntities = append(productEntities, productEntity)
	}

	return productEntities, nil
}

func (pr *ProductRepository) Update(id uint, productEntity *domain.ProductEntity) error {
	// Find the existing product
	var product models.Product
	if err := pr.db.First(&product, id).Error; err != nil {
		return err
	}

	// Update the product fields
	product.Name = productEntity.Name
	product.Description = productEntity.Description
	product.Price = productEntity.Price
	product.SKU = productEntity.SKU
	product.StockQty = productEntity.StockQty
	product.CategoryID = productEntity.CategoryID

	// Save the updated product
	if err := pr.db.Save(&product).Error; err != nil {
		return err
	}

	return nil
}

func (pr *ProductRepository) Delete(id uint) error {
	// Find the product by ID
	var product models.Product
	if err := pr.db.First(&product, id).Error; err != nil {
		return err
	}

	// Delete the product
	if err := pr.db.Delete(&product).Error; err != nil {
		return err
	}

	return nil
}
