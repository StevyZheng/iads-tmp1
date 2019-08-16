package config

import (
	"github.com/spf13/viper"
	"iads/server/internals/pkg/config"
)

func LoadDefaultConfig() *config.Config {
	cfg := &config.Config{}

	cfg.Web.StaticPath = ""
	cfg.Web.Port = 8080
	cfg.Web.ReadTimeout = 0
	cfg.Web.WriteTimeout = 0
	cfg.Web.IdleTimeout = 0

	cfg.Gorm.DBType = "mysql"
	cfg.Gorm.Debug = true
	cfg.MySQL.Host = "127.0.0.1"
	cfg.MySQL.Port = 3306
	cfg.MySQL.User = "root"
	cfg.MySQL.Password = "000000"
	cfg.MySQL.DBName = "iads"
	cfg.MySQL.Parameters = ""

	if cfg.Gorm.DBType == "mysql" {
		cfg.Gorm.DSN = cfg.MySQL.DSN()
	} else if cfg.Gorm.DBType == "sqlite3" {
		cfg.Gorm.DSN = cfg.Sqlite3.DSN()
	}
	return cfg
}

// 加载配置
func LoadConfig(fpath string) (c *config.Config, err error) {
	v := viper.New()
	v.SetConfigFile(fpath)
	v.SetConfigType("yaml")
	if err1 := v.ReadInConfig(); err1 != nil {
		err = err1
		return
	}
	c = &config.Config{}
	c.Web.StaticPath = v.GetString("web.static_path")
	c.Web.Domain = v.GetString("web.domain")
	c.Web.Port = v.GetInt("web.port")
	c.Web.ReadTimeout = v.GetInt("web.read_timeout")
	c.Web.WriteTimeout = v.GetInt("web.write_timeout")
	c.Web.IdleTimeout = v.GetInt("web.idle_timeout")
	c.MySQL.Host = v.GetString("mysql.host")
	c.MySQL.Port = v.GetInt("mysql.port")
	c.MySQL.User = v.GetString("mysql.user")
	c.MySQL.Password = v.GetString("mysql.password")
	c.MySQL.DBName = v.GetString("mysql.db_name")
	c.MySQL.Parameters = v.GetString("mysql.parameters")
	c.Sqlite3.Path = v.GetString("sqlite3.path")
	c.Gorm.Debug = v.GetBool("gorm.debug")
	c.Gorm.DBType = v.GetString("gorm.db_type")
	c.Gorm.MaxLifetime = v.GetInt("gorm.max_lifetime")
	c.Gorm.MaxOpenConns = v.GetInt("gorm.max_open_conns")
	c.Gorm.MaxIdleConns = v.GetInt("gorm.max_idle_conns")
	c.Gorm.TablePrefix = v.GetString("gorm.table_prefix")
	return
}
