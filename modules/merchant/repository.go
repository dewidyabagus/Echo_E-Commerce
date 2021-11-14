package merchant

// ==== Attention, this is a dummy data structure for trader data ====

import "time"

type Merchant struct {
	ID        string    `gorm:"id;type:uuid;primaryKey"`
	Name      string    `gorm:"name;type:varchar(100);index:merchants_name_idx;not null"`
	CreatedAt time.Time `gorm:"created_at;not null"`
	UpdatedAt time.Time `gorm:"updated_at;not null"`
	Deleted   bool      `gorm:"deleted;type:boolean;default:false"`
	DeletedAt time.Time `gorm:"deleted_at"`
}
