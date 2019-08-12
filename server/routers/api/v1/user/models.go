package user

import (
	"github.com/jinzhu/gorm"
	"time"
)

type user struct {
	gorm.Model
	ID         int       `gorm:"AUTO_INCREMENT;primary_key" json:"id"`
	Username   string    `gorm:"type:varchar(64);not null" json:"username"`
	Password   string    `gorm:"type:varchar(64);not null" json:"password"`
	Sex        string    `gorm:"type:varchar(6)" json:"sex"`
	UpdateTime time.Time `json:"update_time"`
	Role       string    `json:"role"`
}
