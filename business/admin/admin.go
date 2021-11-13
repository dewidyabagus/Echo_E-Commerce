package admin

import (
	"crypto/md5"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Admin struct {
	ID        string
	Email     string
	FirstName string
	LastName  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AdminSpec struct {
	Email     string `validate:"required,email"`
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Password  string `validate:"required"`
}

func (a *AdminSpec) toInsert() *Admin {
	password := fmt.Sprintf("%x", md5.Sum([]byte(a.Password)))

	return &Admin{
		ID:        uuid.NewString(),
		Email:     a.Email,
		FirstName: a.FirstName,
		LastName:  a.LastName,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
