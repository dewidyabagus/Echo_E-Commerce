package request

import "RESTful/business/user"

type RequestUser struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	OutletId  string `json:"outlet_id"`
	Password  string `json:"password"`
}

func (r *RequestUser) ToUserAddSpec() *user.UserAddSpec {
	return &user.UserAddSpec{
		Email:     r.Email,
		FirstName: r.FirstName,
		LastName:  r.LastName,
		OutletId:  r.OutletId,
		Password:  r.Password,
	}
}
