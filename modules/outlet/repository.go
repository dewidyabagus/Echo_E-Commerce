package outlet

// === Attention, this is a dummy data structure for data outlets ===

import "time"

type Outlet struct {
	ID         string `gorm:"id;type:uuid;primaryKey"`
	MerchantID string `gorm:"merchant_id;type:uuid;index:outlets_merchant_id_idx;not null"`
	Merchant   Merchant
	Name       string    `gorm:"name;type:varchar(100);index:outlets_name_idx;not null"`
	CreatedAt  time.Time `gorm:"created_at;not null"`
	UpdatedAt  time.Time `gorm:"updated_at;not null"`
	Deleted    bool      `gorm:"deleted;type:boolean;default:false"`
	DeletedAt  time.Time `gorm:"deleted_at"`
}

type Merchant struct {
	ID   string `gorm:"id;type:uuid;primaryKey"`
	Name string `gorm:"name;type:varchar(100);index:merchants_name_idx;not null"`
}
