package migration

import (
	"gorm.io/gorm"

	"RESTful/modules/user"
)

func TableMigration(db *gorm.DB) {
	db.AutoMigrate(&user.User{})
}
