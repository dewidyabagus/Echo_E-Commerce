package admin

import (
	"RESTful/business"
	"RESTful/utils/validator"
)

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) AddNewAdmin(admin *AdminSpec) error {
	if err := validator.GetValidator().Struct(admin); err != nil {
		return business.ErrDataNotSpec
	}

	if _, err := s.repository.VerifyAdminEmail(&admin.Email); err == nil {
		return business.ErrDataConflict

	} else if err != business.ErrDataNotFound {
		return err

	}
	return s.repository.AddNewAdmin(admin.toInsert())
}

func (s *service) GetAdminWithEmailPassword(email, password *string) (id *string, err error) {
	return s.repository.GetAdminWithEmailPassword(email, password)
}

func (s *service) GetAdminById(id *string) (*Admin, error) {
	return s.repository.GetAdminById(id)
}
