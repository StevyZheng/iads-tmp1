package sys

import (
	"github.com/jinzhu/gorm"
	"iads/server/internals/pkg/models/basemodel"
	"iads/server/internals/pkg/models/db"
	"time"
)

// 用户-角色
type UserRole struct {
	basemodel.Model
	UserID uint64 `gorm:"column:admins_id;unique_index:uk_admins_role_admins_id;not null;"` // 管理员ID
	RoleID uint64 `gorm:"column:role_id;unique_index:uk_admins_role_admins_id;not null;"`   // 角色ID
}

// 表名
func (UserRole) TableName() string {
	return TableName("user_role")
}

// 添加前
func (m *UserRole) BeforeCreate(scope *gorm.Scope) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

// 更新前
func (m *UserRole) BeforeUpdate(scope *gorm.Scope) error {
	m.UpdatedAt = time.Now()
	return nil
}

// 分配用户角色
func (UserRole) SetRole(adminsid uint64, roleids []uint64) error {
	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where(&UserRole{UserID: adminsid}).Delete(&UserRole{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if len(roleids) > 0 {
		for _, rid := range roleids {
			rm := new(UserRole)
			rm.RoleID = rid
			rm.UserID = adminsid
			if err := tx.Create(rm).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	return tx.Commit().Error
}
