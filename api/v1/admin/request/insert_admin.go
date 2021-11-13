package request

import (
	"RESTful/business/admin"
)

type RequestAdmin struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

func (r *RequestAdmin) ToAdminAddSpec() *admin.AdminSpec {
	return &admin.AdminSpec{
		Email:     r.Email,
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Password:  r.Password,
	}
}
