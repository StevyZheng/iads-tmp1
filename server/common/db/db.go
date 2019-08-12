package db

import (
	"github.com/jinzhu/gorm"
	//"iads/server/common/config"
)

var DB *gorm.DB

type Connection struct {
	conn *gorm.DB
}

func NewConnection() (conn *Connection) {

}
