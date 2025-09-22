package models

import (
	"context"

	"github.com/akasection/durianpay-cs-dashboard/backend/pkg/util"
	"github.com/akasection/durianpay-cs-dashboard/backend/services"
	"gorm.io/gorm"
)

type User struct {
	ID        uint     `gorm:"primaryKey" json:"id"`
	Username  string   `gorm:"uniqueIndex;size:31" json:"username"`
	HPassword string   `json:"password"`
	Roles     []string `json:"roles"`
}

func CheckCredentials(username, password string) bool {
	ctx := context.Background()
	result, err := gorm.G[User](services.DB).Where("username = ?", username).First(ctx)
	// Match the password and hash
	if err != nil && err != gorm.ErrRecordNotFound {
		return false
	}
	isMatch := util.MatchPassword(password, result.HPassword)

	return isMatch
}
