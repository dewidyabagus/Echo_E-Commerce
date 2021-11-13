package response

import (
	"RESTful/business/user"
	"time"
)

type UserDetail struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	OutletId  string    `json:"outlet_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetUserResponseDetail(user *user.User) *UserDetail {
	return &UserDetail{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		OutletId:  user.OutletId,
		UpdatedAt: user.UpdatedAt,
	}
}

func GetAllUserResponseDetail(users *[]user.User) *[]UserDetail {
	var response = []UserDetail{}

	for _, user := range *users {
		response = append(response, *GetUserResponseDetail(&user))
	}

	return &response
}
