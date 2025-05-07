package dto

import (
	"fmt"
	"net/http"
)

type CreateProductRequest struct {
	Name         string  `json:"name" binding:"required"`
	Desc         string  `json:"desc" binding:"required"`
	SKU          string  `json:"sku" binding:"required"`
	CategoryName string  `json:"categoryName" binding:"required"`
	Price        float64 `json:"price" binding:"required"`
	StockQty     uint    `json:"stock_qty" binding:"required"`
}

type CreateProductResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type FetchProductResponse struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Desc     string  `json:"desc"`
	Price    float64 `json:"price"`
	SKU      string  `json:"sku"`
	StockQty uint    `json:"stock_qty"`
	Category string  `json:"category"`
}

type UpdateProductResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type DeleteProductResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type CreateOrUpdateProductDto struct {
	Name         *string
	Desc         *string
	SKU          *string
	CategoryName *string
	Price        *float64
	StockQty     *uint
}

func NewCreateProductResponse(name string) *CreateProductResponse {
	return &CreateProductResponse{
		StatusCode: http.StatusCreated,
		Message:    fmt.Sprintf("%s created successfully", name),
	}
}

func NewUpdateProductResponse(name string) *UpdateProductResponse {
	return &UpdateProductResponse{
		StatusCode: http.StatusOK,
		Message:    fmt.Sprintf("%s updated successfully", name),
	}
}

func NewDeleteProductResponse(name string) *DeleteProductResponse {
	return &DeleteProductResponse{
		StatusCode: http.StatusOK,
		Message:    "product deleted successfully",
	}
}
