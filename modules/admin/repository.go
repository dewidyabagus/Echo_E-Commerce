package admin

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"RESTful/business"
	"RESTful/business/admin"
)

type Admin struct {
	ID        string    `gorm:"id;type:uuid;primaryKey"`
	Email     string    `gorm:"email;type:varchar(100);index:admins_email_idx;not null"`
	FirstName string    `gorm:"first_name;type:varchar(100);not null"`
	LastName  string    `gorm:"last_name;type:varchar(100);not null"`
	Password  string    `gorm:"password;type:varchar(32);not null"`
	CreatedAt time.Time `gorm:"created_at;type:timestamp;not null"`
	UpdatedAt time.Time `gorm:"updated_at;type:timestamp;not null"`
	Deleted   bool      `gorm:"deleted;type:boolean;default:false"`
	DeletedAt time.Time `gorm:"deleted_at;type:timestamp"`
}

func (a *Admin) toBusinessAdmin() *admin.Admin {
	return &admin.Admin{
		ID:        a.ID,
		Email:     a.Email,
		FirstName: a.FirstName,
		LastName:  a.LastName,
		Password:  a.Password,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}

type AdminEmail struct {
	Email string `gorm:"email"`
}

type AdminID struct {
	ID string `gorm:"id"`
}

type Repository struct {
	DB *gorm.DB
}

func toInsertAdmin(admin *admin.Admin) *Admin {
	return &Admin{
		ID:        admin.ID,
		Email:     admin.Email,
		FirstName: admin.FirstName,
		LastName:  admin.LastName,
		Password:  admin.Password,
		CreatedAt: admin.CreatedAt,
		UpdatedAt: admin.UpdatedAt,
	}
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) AddNewAdmin(admin *admin.Admin) error {
	return r.DB.Create(toInsertAdmin(admin)).Error
}

func (r *Repository) VerifyAdminEmail(email *string) (*string, error) {
	verify := new(AdminEmail)

	err := r.DB.Model(&Admin{}).First(verify, "email = ? and deleted = false", email).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, business.ErrDataNotFound
		}
		return nil, err
	}

	return &verify.Email, nil
}

func (r *Repository) GetAdminWithEmailPassword(email *string, password *string) (*string, error) {
	result := new(AdminID)

	err := r.DB.Model(&Admin{}).First(result, "email = ? and password = md5(?) and deleted = false", email, password).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, business.ErrUnauthorized
		}
		return nil, err
	}

	return &result.ID, nil
}

func (r *Repository) GetAdminById(id *string) (*admin.Admin, error) {
	result := new(Admin)

	err := r.DB.First(result, "id = ? and deleted = false", id).Error
	if err != nil {
		return nil, err
	}

	return result.toBusinessAdmin(), nil
}
