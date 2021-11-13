package auth

import (
	"time"

	"github.com/golang-jwt/jwt"

	"RESTful/business"
	"RESTful/business/user"
	"RESTful/config"
	"RESTful/utils/validator"
)

type service struct {
	userService user.Service
}

func NewService(userService user.Service) Service {
	return &service{userService}
}

func (s *service) UserLoginWithEmailPassword(user *LoginSpec) (token *string, err error) {
	err = validator.GetValidator().Struct(user)
	if err != nil {
		return nil, business.ErrUnauthorized
	}

	id, err := s.userService.GetUserWithEmailPassword(&user.Email, &user.Password)
	if err != nil {
		return token, err
	}

	genToken, err := createToken(id)
	if err != nil {
		return token, err
	}

	return &genToken, nil
}

func createToken(id *string) (token string, err error) {
	eJWT := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":  *id,
			"exp": time.Now().Add(time.Minute * 60).Unix(),
		},
	)

	token, err = eJWT.SignedString([]byte(config.GetJWTSecretKey()))
	if err != nil {
		return "", err
	}

	return token, err
}
