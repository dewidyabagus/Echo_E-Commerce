package product

type Service interface {
	// Get all product
	FindAllProduct(requester *string) (*[]Product, error)

	// Get product detail by id
	// FindProductById(id, requester *string) (*Product, error)

	// Add new product
	AddNewProduct(product *ProductSpec, addter *string) error

	// Modify Information Product
	// UpdateProduct(id *string, product *ProductSpec, modifier *string) error

	// Delete product
	// DeleteProduct(id, deleter *string) error
}

type Repository interface {
	// Get all product
	FindAllProduct(merchant_id *string) (*[]Product, error)

	// Get product detail by id
	// FindProductById(id *string) (*Product, error)

	// Get product by sku
	FindProductBySKU(merchant_id, sku *string) (id *string, err error)

	// Add new product
	AddNewProduct(product *Product) error

	// Modify Information Product
	// UpdateProduct(id *string, product *Product) error

	// Delete product
	// DeleteProduct(id *string) error
}
