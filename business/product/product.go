package product

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID           string
	MerchantID   string
	MerchantName string
	SKU          string
	Name         string
	CategoryName string
	Description  string
	Unit         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type ProductSpec struct {
	SKU          string `validate:"required"`
	Name         string `validate:"required"`
	CategoryName string `validate:"required"`
	Description  string `validate:"required"`
	Unit         string `validate:"required"`
}

func (p *ProductSpec) toInsertProduct(merchatId *string) *Product {
	return &Product{
		ID:           uuid.NewString(),
		MerchantID:   *merchatId,
		SKU:          p.SKU,
		Name:         p.Name,
		CategoryName: p.CategoryName,
		Description:  p.Description,
		Unit:         p.Unit,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

// func (p *ProductSpec) toUpdateProduct() *Product {
// 	return &Product{
// 		SKU:          p.SKU,
// 		Name:         p.Name,
// 		CategoryName: p.CategoryName,
// 		Description:  p.Description,
// 		Unit:         p.Unit,
// 		UpdatedAt:    time.Now(),
// 	}
// }
