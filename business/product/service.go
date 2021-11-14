package product

import (
	"RESTful/business"
	"RESTful/business/user"
	"RESTful/utils/validator"
)

type service struct {
	repository  Repository
	userService user.Service
}

func NewService(repository Repository, userService user.Service) Service {
	return &service{repository, userService}
}

// Add new product
func (s *service) AddNewProduct(product *ProductSpec, addter *string) error {
	mapUser, err := s.userService.GetUserByIdFromService(addter)
	if err != nil {
		return business.ErrUnauthorized
	}
	merchant_id, _ := mapUser["merchant_id"].(string)

	if err := validator.GetValidator().Struct(product); err != nil {
		return business.ErrDataNotSpec
	}

	if _, err := s.repository.FindProductBySKU(&merchant_id, &product.SKU); err == nil {
		return business.ErrDataConflict
	}

	return s.repository.AddNewProduct(product.toInsertProduct(&merchant_id))
}

// Get all product
func (s *service) FindAllProduct(requester *string) (*[]Product, error) {
	mapUser, err := s.userService.GetUserByIdFromService(requester)
	if err != nil {
		return nil, business.ErrUnauthorized
	}
	merchant_id, _ := mapUser["merchant_id"].(string)

	return s.repository.FindAllProduct(&merchant_id)
}

// Get product detail by id
// func (s *service) FindProductById(id, requester *string) (*Product, error) {
// 	if _, err := s.userService.GetUserByIdFromService(requester); err != nil {
// 		return nil, business.ErrUnauthorized
// 	}
// }

// Modify Information Product
// func (s *service) UpdateProduct(id *string, product *ProductSpec, modifier *string) error {
// 	if _, err := s.userService.GetUserByIdFromService(modifier); err != nil {
// 		return nil
// 	}
// }

// Delete product
// func (s *service) DeleteProduct(id, deleter *string) error {
// 	if _, err := s.userService.GetUserByIdFromService(deleter); err != nil {
// 		return err
// 	}
// }
