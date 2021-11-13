package auth

type Service interface {
	// Login User
	UserLoginWithEmailPassword(user *LoginSpec) (token *string, err error)

	// Login Admin
	AdminLoginWithEmailPassword(admin *LoginSpec) (token *string, err error)
}
