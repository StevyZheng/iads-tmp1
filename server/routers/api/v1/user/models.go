package user

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID         int       `gorm:"AUTO_INCREMENT;primary_key" json:"id"`
	Username   string    `gorm:"type:varchar(64);not null" json:"username"`
	Password   string    `gorm:"type:varchar(64);not null" json:"password"`
	Sex        string    `gorm:"type:varchar(6)" json:"sex"`
	Email      string    `gorm:"type:varchar(128)" json:"email"`
	Phone      string    `gorm:"type:varchar(12)" json:"phone"`
	UpdateTime time.Time `json:"update_time"`
	Role       string    `json:"role"`
}
