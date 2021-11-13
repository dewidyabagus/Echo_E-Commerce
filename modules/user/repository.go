package user

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"RESTful/business"
	"RESTful/business/user"
)

type User struct {
	ID        string    `gorm:"id;type:uuid;primaryKey;index:users_id_idx"`
	Email     string    `gorm:"email;type:varchar(100);index:users_email_idx;not null"`
	FirstName string    `gorm:"first_name;type:varchar(100);not null"`
	LastName  string    `gorm:"last_name;type:varchar(100);not null"`
	Password  string    `gorm:"password;type:varchar(32);not null"`
	OutletId  string    `gorm:"outlet_id;type:uuid;not null"`
	CreatedAt time.Time `gorm:"created_at;type:timestamp;not null"`
	UpdatedAt time.Time `gorm:"updated_at;type:timestamp;not null"`
	DeletedAt time.Time `gorm:"index"`
}

func (u *User) toBusinessUser() *user.User {
	return &user.User{
		ID:        u.ID,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		OutletId:  u.OutletId,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

type UserLogin struct {
	ID string `gorm:"id"`
}

type Repository struct {
	DB *gorm.DB
}

func toUserInsert(u *user.User) *User {
	return &User{
		ID:        u.ID,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Password:  u.Password,
		OutletId:  u.OutletId,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) AddNewUser(user *user.User) error {
	testUser := new(User)

	err := r.DB.First(testUser, "email = ?", user.Email).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		if err != nil {
			return err
		}
		return business.ErrDataConflict
	}

	return r.DB.Create(toUserInsert(user)).Error
}

func (r *Repository) GetUserWithEmailPassword(email *string, password *string) (id *string, err error) {
	var login = new(UserLogin)

	err = r.DB.Model(&User{}).First(login, "email = ? and password = md5(?)", email, password).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return id, business.ErrUnauthorized
		}
		return id, err
	}

	return &login.ID, nil
}

func (r *Repository) FindUserByUserId(id *string) (*user.User, error) {
	var user = new(User)

	if err := r.DB.First(user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return user.toBusinessUser(), nil
}
