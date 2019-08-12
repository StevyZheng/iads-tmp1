package db

import (
	"github.com/jinzhu/gorm"
	"iads/server/common/config"
	//"iads/server/common/config"
)

var DB *gorm.DB

type Connection struct {
	DB *gorm.DB
}

func NewConnection() (conn *Connection) {
	var err error
	conn = new(Connection)
	cfg := config.LoadDefaultConfig()
	conn.DB, err = gorm.Open(cfg.Gorm.DBType, cfg.Gorm.DSN)
}
