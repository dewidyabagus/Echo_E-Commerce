package admin

type Service interface {
	// AddNewAdmin for new admin registration, returns nil on success.
	AddNewAdmin(admin *AdminSpec) error

	// GetAdminWithEmailPassword using admin user login, returns token and nil on success
	GetAdminWithEmailPassword(email *string, password *string) (id *string, err error)

	// GetAdminById
	GetAdminById(id *string) (*Admin, error)
}

type Repository interface {
	// AddNewAdmin for new admin registration, returns nil on success.
	AddNewAdmin(admin *Admin) error

	// GetAdminWithEmailPassword using admin user login, returns token and nil on success
	GetAdminWithEmailPassword(email *string, password *string) (id *string, err error)

	// VerifyAdminEmail verify email already exists
	VerifyAdminEmail(email *string) (*string, error)

	// GetAdminById
	GetAdminById(id *string) (*Admin, error)
}
