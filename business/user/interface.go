package user

import "time"

type Service interface {
	// Find all user
	FindAllUser(requeser *string) (*[]User, error)

	// Find user by user id
	FindUserByUserId(id, requeser *string) (*User, error)

	// Get user with email and password. for user login
	GetUserWithEmailPassword(email, password *string) (id *string, err error)

	// Add new user
	AddNewUser(user *UserAddSpec, addter *string) error

	// Update user by id
	UpdateUser(id *string, user *UserEditSpec, modifier *string) error

	// Delete user by id
	DeleteUser(id, deleter *string) error
}

type Repository interface {
	// Find all user
	FindAllUser() (*[]User, error)

	// Find user by user id
	FindUserByUserId(id *string) (*User, error)

	// Get user with email and password. for user login
	GetUserWithEmailPassword(email, password *string) (id *string, err error)

	// Verify user email
	VerifyUserEmail(email *string) (id *string, err error)

	// Add new user
	AddNewUser(user *User) error

	// Update user by id
	UpdateUser(id *string, user *User) error

	// Delete user by id
	DeleteUser(id *string, deletedAt time.Time) error
}
