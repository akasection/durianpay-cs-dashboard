package bootstrap

import (
	"gorm.io/gorm"

	"github.com/akasection/durianpay-cs-dashboard/backend/models"
)

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&models.Role{})
	db.AutoMigrate(&models.Permission{})
	db.AutoMigrate(&models.UserRole{})
	db.AutoMigrate(&models.RolePermission{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Payment{})
}
