package request

import "RESTful/business/user"

type RequestUser struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	OutletID  string `json:"outlet_id"`
	Password  string `json:"password"`
}

func (r *RequestUser) ToUserAddSpec() *user.UserAddSpec {
	return &user.UserAddSpec{
		Email:     r.Email,
		FirstName: r.FirstName,
		LastName:  r.LastName,
		OutletID:  r.OutletID,
		Password:  r.Password,
	}
}

func (r *RequestUser) ToUserUpateSpec() *user.UserEditSpec {
	return &user.UserEditSpec{
		Email:     r.Email,
		FirstName: r.FirstName,
		LastName:  r.LastName,
		OutletID:  r.OutletID,
	}
}
