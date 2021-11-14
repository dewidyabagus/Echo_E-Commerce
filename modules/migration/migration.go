package migration

import (
	"gorm.io/gorm"

	"RESTful/modules/admin"
	"RESTful/modules/merchant"
	"RESTful/modules/outlet"
	"RESTful/modules/product"
	"RESTful/modules/user"
)

func TableMigration(db *gorm.DB) {
	db.AutoMigrate(&admin.Admin{}, &merchant.Merchant{}, &outlet.Outlet{}, &user.User{}, &product.Product{})
}
