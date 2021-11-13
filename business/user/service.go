package user

import (
	"RESTful/business"
	"RESTful/business/admin"
	"RESTful/utils/validator"
	"time"
)

type service struct {
	repository Repository
	admin      admin.Service
}

func NewService(repository Repository, admin admin.Service) Service {
	return &service{repository, admin}
}

func (s *service) AddNewUser(user *UserAddSpec, addter *string) error {
	if _, err := s.admin.GetAdminById(addter); err != nil {
		return business.ErrUnauthorized
	}

	err := validator.GetValidator().Struct(user)
	if err != nil {
		return business.ErrDataNotSpec
	}

	_, err = s.repository.VerifyUserEmail(&user.Email)
	if err == nil {
		return business.ErrDataConflict

	} else if err != business.ErrDataNotFound {
		return err

	}

	return s.repository.AddNewUser(user.toInsertUser())
}

func (s *service) GetUserWithEmailPassword(email *string, password *string) (id *string, err error) {
	return s.repository.GetUserWithEmailPassword(email, password)
}

func (s *service) FindAllUser(requeser *string) (*[]User, error) {
	if _, err := s.admin.GetAdminById(requeser); err != nil {
		return nil, business.ErrUnauthorized
	}

	return s.repository.FindAllUser()
}

func (s *service) FindUserByUserId(id, requeser *string) (*User, error) {
	if _, err := s.admin.GetAdminById(requeser); err != nil {
		return nil, business.ErrUnauthorized
	}

	return s.repository.FindUserByUserId(id)
}

func (s *service) UpdateUser(id *string, user *UserEditSpec, modifier *string) error {
	if _, err := s.admin.GetAdminById(modifier); err != nil {
		return business.ErrUnauthorized
	}

	if err := validator.GetValidator().Struct(user); err != nil {
		return business.ErrDataNotSpec
	}

	if _, err := s.repository.FindUserByUserId(id); err != nil {
		return err
	}

	userId, err := s.repository.VerifyUserEmail(&user.Email)
	if err == nil {
		if *userId != *id {
			return business.ErrDataConflict
		}

	} else if err != business.ErrDataNotFound {
		return err

	}

	return s.repository.UpdateUser(id, user.toUpdateUser())
}

func (s *service) DeleteUser(id, deleter *string) error {
	if _, err := s.admin.GetAdminById(deleter); err != nil {
		return business.ErrUnauthorized
	}

	if _, err := s.repository.FindUserByUserId(id); err != nil {
		return err
	}

	return s.repository.DeleteUser(id, time.Now())
}
