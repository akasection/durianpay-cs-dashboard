package models

import (
	"context"

	"github.com/akasection/durianpay-cs-dashboard/backend/pkg/util"
	"github.com/akasection/durianpay-cs-dashboard/backend/services"
	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey;AUTO_INCREMENT" json:"id"`
	Username  string `gorm:"uniqueIndex;size:31;type:varchar(32);not null" json:"username"`
	HPassword string `gorm:"type:varchar(72);not null" json:"password"`
}

func CheckCredentials(username, password string) bool {
	result, err := GetUserByUsername(username)
	// Match the password and hash
	if err != nil && err != gorm.ErrRecordNotFound {
		return false
	}
	isMatch := util.MatchPassword(result.HPassword, password)

	return isMatch
}

func GetUserByUsername(username string) (User, error) {
	ctx := context.Background()
	result, err := gorm.G[User](services.DB).Where("username = ?", username).First(ctx)
	if err != nil {
		return User{}, err
	}
	return result, nil
}

func GetUserRoles(username string) ([]string, error) {
	user, err := GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	var result []map[string]interface{}

	// Gorm generics way too funky on Left Join table. Use tradi instead
	query := services.DB.WithContext(ctx).Table("user_role").
		Select("roles.name").
		Joins("JOIN roles ON user_role.role_id = roles.id").
		Where("user_role.user_id = ?", user.ID)

	qErr := query.Scan(&result).Error
	if qErr != nil {
		return nil, qErr
	}

	var roles []string
	for _, ur := range result {
		roles = append(roles, ur["name"].(string))
	}

	return roles, nil
}

func (User) TableName() string {
	return "users"
}
