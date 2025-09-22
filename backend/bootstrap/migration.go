package bootstrap

import (
	"gorm.io/gorm"

	PaymentModel "github.com/akasection/durianpay-cs-dashboard/backend/models"
	UserModel "github.com/akasection/durianpay-cs-dashboard/backend/models"
)

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&UserModel.User{}, &PaymentModel.Payment{})
}
