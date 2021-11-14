package product

import (
	"time"

	"gorm.io/gorm"

	"RESTful/business/product"
)

type Product struct {
	ID           string `gorm:"id;type:uuid;primaryKey"`
	MerchantID   string `gorm:"merchant_id;uuid;index:products_merchant_idx;not null"`
	Merchant     Merchant
	SKU          string    `gorm:"sku;type:varchar(25);index:products_merchant_idx;not null"`
	Name         string    `gorm:"name;type:varchar(100);not null"`
	CategoryName string    `gorm:"category_name;varchar(100);not null"`
	Description  string    `gorm:"description;type:varchar(200);not null"`
	Unit         string    `gorm:"unit;type:varchar(20);not null"`
	CreatedAt    time.Time `gorm:"created_at;not null"`
	UpdatedAt    time.Time `gorm:"updated_at;not null"`
	Deleted      bool      `gorm:"deleted;type:boolean;default:false"`
	DeletedAt    time.Time `gorm:"deleted_at"`
}

type Merchant struct {
	ID   string `gorm:"id;type:uuid;primaryKey"`
	Name string `gorm:"name;type:varchar(100);not null"`
}

type ProductId struct {
	ID string `gorm:"id"`
}

type Repository struct {
	DB *gorm.DB
}

func (p *Product) toBusinessProduct() *product.Product {
	return &product.Product{
		ID:           p.ID,
		MerchantID:   p.MerchantID,
		MerchantName: p.Merchant.Name,
		SKU:          p.SKU,
		Name:         p.Name,
		CategoryName: p.CategoryName,
		Description:  p.Description,
		Unit:         p.Unit,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
	}
}

func toInsertProduct(p *product.Product) *Product {
	return &Product{
		ID:           p.ID,
		MerchantID:   p.MerchantID,
		SKU:          p.SKU,
		Name:         p.Name,
		CategoryName: p.CategoryName,
		Description:  p.Description,
		Unit:         p.Unit,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
	}
}

func toAllBusinessProduct(products *[]Product) *[]product.Product {
	var response = []product.Product{}

	for _, product := range *products {
		response = append(response, *product.toBusinessProduct())
	}

	return &response
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) AddNewProduct(product *product.Product) error {
	return r.DB.Create(toInsertProduct(product)).Error
}

func (r *Repository) FindProductBySKU(merchant_id, sku *string) (id *string, err error) {
	var productId = new(ProductId)

	err = r.DB.Model(&Product{}).
		First(productId, "deleted = false and merchant_id = ? and sku = ?", merchant_id, sku).Error
	if err != nil {
		return nil, err
	}

	return &productId.ID, nil
}

func (r *Repository) FindAllProduct(merchant_id *string) (*[]product.Product, error) {
	var products = new([]Product)

	err := r.DB.Preload("Merchant").
		Where("deleted = false").
		Where("merchant_id = ?", merchant_id).Find(products).Error
	if err != nil {
		return nil, err
	}

	return toAllBusinessProduct(products), nil
}
