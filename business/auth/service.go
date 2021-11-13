package auth

import (
	"time"

	"github.com/golang-jwt/jwt"

	"RESTful/business"
	"RESTful/business/admin"
	"RESTful/business/user"
	"RESTful/config"
	"RESTful/utils/validator"
)

type service struct {
	userService  user.Service
	adminService admin.Service
}

func NewService(userService user.Service, adminService admin.Service) Service {
	return &service{userService, adminService}
}

func (s *service) UserLoginWithEmailPassword(user *LoginSpec) (token *string, err error) {
	err = validator.GetValidator().Struct(user)
	if err != nil {
		return nil, business.ErrUnauthorized
	}

	id, err := s.userService.GetUserWithEmailPassword(&user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return createToken(id)
}

func (s *service) AdminLoginWithEmailPassword(user *LoginSpec) (token *string, err error) {
	err = validator.GetValidator().Struct(user)
	if err != nil {
		return nil, business.ErrUnauthorized
	}

	id, err := s.adminService.GetAdminWithEmailPassword(&user.Email, &user.Password)
	if err != nil {
		return token, err
	}

	return createToken(id)
}

func createToken(id *string) (*string, error) {
	eJWT := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":  *id,
			"exp": time.Now().Add(time.Minute * 60).Unix(),
		},
	)

	token, err := eJWT.SignedString([]byte(config.GetJWTSecretKey()))
	if err != nil {
		return nil, err
	}

	return &token, err
}
