package response

import (
	"time"

	"RESTful/business/product"
)

type Product struct {
	ID           string    `json:"id"`
	SKU          string    `json:"sku"`
	Name         string    `json:"name"`
	CategoryName string    `json:"category_name"`
	Description  string    `json:"description"`
	Unit         string    `json:"unit"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func GetProductDetail(product *product.Product) *Product {
	return &Product{
		ID:           product.ID,
		SKU:          product.SKU,
		Name:         product.Name,
		CategoryName: product.CategoryName,
		Description:  product.Description,
		Unit:         product.Unit,
		UpdatedAt:    product.UpdatedAt,
	}
}

func GetAllProductDetail(products *[]product.Product) *[]Product {
	var response = []Product{}

	for _, product := range *products {
		response = append(response, *GetProductDetail(&product))
	}

	return &response
}
