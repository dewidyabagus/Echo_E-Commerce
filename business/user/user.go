package user

import (
	"crypto/md5"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        string
	Email     string
	FirstName string
	LastName  string
	OutletID  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserAddSpec struct {
	Email     string `validate:"required,email"`
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	OutletID  string `validate:"required"`
	Password  string `validate:"required"`
}

func (u *UserAddSpec) toInsertUser() *User {
	password := fmt.Sprintf("%x", md5.Sum([]byte(u.Password)))

	return &User{
		ID:        uuid.NewString(),
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		OutletID:  u.OutletID,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type UserEditSpec struct {
	Email     string `validate:"required,email"`
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	OutletID  string `validate:"required"`
}

func (u *UserEditSpec) toUpdateUser() *User {
	return &User{
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		OutletID:  u.OutletID,
		UpdatedAt: time.Now(),
	}
}
