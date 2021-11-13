package user

import (
	"RESTful/business"
	"RESTful/business/admin"
	"RESTful/utils/validator"
)

type service struct {
	repository Repository
	admin      admin.Service
}

func NewService(repository Repository, admin admin.Service) Service {
	return &service{repository, admin}
}

func (s *service) AddNewUser(user *UserAddSpec, addter *string) error {
	err := validator.GetValidator().Struct(user)
	if err != nil {
		return business.ErrDataNotSpec
	}

	if _, err := s.admin.GetAdminById(addter); err != nil {
		return business.ErrUnauthorized
	}

	return s.repository.AddNewUser(user.toInsertUser())
}

func (s *service) GetUserWithEmailPassword(email *string, password *string) (id *string, err error) {
	return s.repository.GetUserWithEmailPassword(email, password)
}

func (s *service) FindUserByUserId(id *string) (*User, error) {
	return s.repository.FindUserByUserId(id)
}
