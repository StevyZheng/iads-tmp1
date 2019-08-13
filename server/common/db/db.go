package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"iads/server/common"
	"iads/server/common/config"
	"iads/server/routers/api/v1/user"
	"log"
	"os"
	"time"
	//"iads/server/common/config"
)

var conn Connection

type Connection struct {
	DB *gorm.DB
}

func NewConnection() (conn *Connection) {
	var err error
	conn = new(Connection)
	cfg := config.LoadDefaultConfig()
	conn.DB, err = gorm.Open(cfg.Gorm.DBType, cfg.Gorm.DSN)
	common.CheckErr(err)
	conn.DB.SingularTable(true)
	if cfg.Gorm.Debug {
		conn.DB.LogMode(true)
		conn.DB.SetLogger(log.New(os.Stdout, "\n", 0))
	}
	conn.DB.DB().SetMaxIdleConns(cfg.Gorm.MaxIdleConns)
	conn.DB.DB().SetMaxOpenConns(cfg.Gorm.MaxOpenConns)
	conn.DB.DB().SetConnMaxLifetime(time.Duration(cfg.Gorm.MaxLifetime) * time.Second)
	return conn
}

func (c *Connection) Migration() {
	fmt.Println(c.DB.AutoMigrate(new(user.User)).Error)
}
