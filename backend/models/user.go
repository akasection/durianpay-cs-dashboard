package models

import (
	"context"

	"github.com/akasection/durianpay-cs-dashboard/backend/pkg/util"
	"github.com/akasection/durianpay-cs-dashboard/backend/services"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Username  string `gorm:"uniqueIndex;size:31" json:"username"`
	HPassword string `json:"password"`
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

func GetUserByUsername(username string) (*User, error) {
	ctx := context.Background()
	result, err := gorm.G[User](services.DB).Where("username = ?", username).First(ctx)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func ListUserPermissions(username string) ([]string, error) {
	user, err := GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	result, err := gorm.G[UserRole](services.DB).Where("user_id = ?", user.ID).Find(ctx)
	if err != nil {
		return nil, err
	}

	var roleIds []uint
	for _, ur := range result {
		roleIds = append(roleIds, ur.RoleId)
	}

	return ListRolePermissionsByRoleIds(roleIds)
}

func ListRolePermissionsByRoleIds(roleIds []uint) ([]string, error) {
	ctx := context.Background()
	var result []string
	permissions, err := gorm.G[Permission](services.DB).
		Select("name").
		Joins(
			clause.LeftJoin.AssociationFrom("role_permission", gorm.G[RolePermission](services.DB)).As("t"),
			func(db gorm.JoinBuilder, joinTable, curTable clause.Table) error {
				db.Where("?.permission_id = ?.id", joinTable, curTable)
				return nil
			},
		).
		// Joins("left join role_permission ON role_permission.permission_id = permission.id").
		Where("role_permission.role_id IN (?)", roleIds).
		Find(ctx)
	if err != nil {
		return nil, err
	}

	for _, rp := range permissions {
		result = append(result, rp.Name)
	}

	return result, nil
}

func (User) TableName() string {
	return "users"
}
