package user

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"RESTful/business"
	"RESTful/business/user"
)

type User struct {
	ID        string `gorm:"id;type:uuid;primaryKey;index:users_id_idx"`
	Email     string `gorm:"email;type:varchar(100);index:users_email_idx;not null"`
	FirstName string `gorm:"first_name;type:varchar(100);not null"`
	LastName  string `gorm:"last_name;type:varchar(100);not null"`
	Password  string `gorm:"password;type:varchar(32);not null"`
	OutletID  string `gorm:"outlet_id;type:uuid;index:users_outlet_id_idx;not null"`
	Outlet    Outlet
	CreatedAt time.Time `gorm:"created_at;type:timestamp;not null"`
	UpdatedAt time.Time `gorm:"updated_at;type:timestamp;not null"`
	Deleted   bool      `gorm:"deleted;type:boolean;default:false"`
	DeletedAt time.Time `gorm:"deleted_at;type:timestamp"`
}

type Outlet struct {
	ID         string `gorm:"id;type:uuid;primaryKey"`
	MerchantID string `gorm:"merchant_id;type:uuid;index:outlets_merchant_id_idx;not null"`
	Name       string `gorm:"name;type:varchar(100);index:outlets_name_idx;not null"`
}

func (u *User) toBusinessUser() *user.User {
	return &user.User{
		ID:        u.ID,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		OutletID:  u.OutletID,
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
		OutletID:  u.OutletID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func toUpdateUser(u *user.User) *User {
	return &User{
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Password:  u.Password,
		OutletID:  u.OutletID,
		UpdatedAt: u.UpdatedAt,
	}
}

func toAllBusinessUser(users *[]User) *[]user.User {
	var response = []user.User{}

	for _, user := range *users {
		response = append(response, *user.toBusinessUser())
	}

	return &response
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) AddNewUser(user *user.User) error {
	return r.DB.Create(toUserInsert(user)).Error
}

func (r *Repository) GetUserWithEmailPassword(email *string, password *string) (id *string, err error) {
	var login = new(UserLogin)

	err = r.DB.Model(&User{}).First(login, "deleted = false and email = ? and password = md5(?)", email, password).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return id, business.ErrUnauthorized
		}
		return id, err
	}

	return &login.ID, nil
}

func (r *Repository) VerifyUserEmail(email *string) (id *string, err error) {
	verify := new(UserLogin)

	err = r.DB.Model(&User{}).First(verify, "deleted = false and email = ?", email).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, business.ErrDataNotFound
		}
		return nil, err
	}

	return &verify.ID, nil
}

func (r *Repository) FindAllUser() (*[]user.User, error) {
	var users = new([]User)

	err := r.DB.Where("deleted = false").Find(users).Order("outlet_id asc").Error
	if err != nil {
		return nil, err
	}

	return toAllBusinessUser(users), nil
}

func (r *Repository) FindUserByUserId(id *string) (*user.User, error) {
	var user = new(User)

	if err := r.DB.First(user, "deleted = false and id = ?", id).Error; err != nil {
		return nil, err
	}

	return user.toBusinessUser(), nil
}

func (r *Repository) UpdateUser(id *string, user *user.User) error {
	return r.DB.Model(&User{}).Where("id = ?", id).Updates(toUpdateUser(user)).Error
}

func (r *Repository) DeleteUser(id *string, deletedAt time.Time) error {
	return r.DB.Model(&User{}).Where("id = ?", id).Updates(&User{Deleted: true, DeletedAt: deletedAt}).Error
}
