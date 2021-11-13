package response

import "RESTful/business/user"

type UserDetail struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	OutletId  string `json:"outlet_id"`
}

func GetUserResponseDetail(user *user.User) *UserDetail {
	return &UserDetail{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		OutletId:  user.OutletId,
	}
}
