package request

import (
	"RESTful/business/product"
)

type RequestProduct struct {
	SKU          string `json:"sku"`
	Name         string `json:"name"`
	CategoryName string `json:"category_name"`
	Description  string `json:"description"`
	Unit         string `json:"unit"`
}

func (r *RequestProduct) ToProductSpec() *product.ProductSpec {
	return &product.ProductSpec{
		SKU:          r.SKU,
		Name:         r.Name,
		CategoryName: r.CategoryName,
		Description:  r.Description,
		Unit:         r.Unit,
	}
}
