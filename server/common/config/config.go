package config

import "fmt"

func LoadDefaultConfig() Config {
	cfg := Config{}
	cfg.Gorm.DBType = "mysql"
	cfg.Gorm.Debug = true
	cfg.MySQL.Host = "127.0.0.1"
	cfg.MySQL.Port = 3306
	cfg.MySQL.User = "root"
	cfg.MySQL.Password = "000000"
	cfg.MySQL.DBName = "iads"
	if cfg.Gorm.DBType == "mysql" {
		cfg.Gorm.DSN = cfg.MySQL.DSN()
	} else if cfg.Gorm.DBType == "sqlite3" {
		cfg.Gorm.DSN = cfg.Sqlite3.DSN()
	}
	return cfg
}

// Config 配置参数
type Config struct {
	Web     Web
	Gorm    Gorm
	MySQL   MySQL
	Sqlite3 Sqlite3
}

// 站点配置参数
type Web struct {
	Domain       string
	StaticPath   string
	Port         int
	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int
}

// Gorm gorm配置参数
type Gorm struct {
	Debug        bool
	DBType       string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
	TablePrefix  string
	DSN          string
}

// MySQL mysql配置参数
type MySQL struct {
	Host       string
	Port       int
	User       string
	Password   string
	DBName     string
	Parameters string
}

// MySQL 数据库连接串
func (a MySQL) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		a.User, a.Password, a.Host, a.Port, a.DBName, a.Parameters)
}

// Sqlite3 配置参数
type Sqlite3 struct {
	Path string
}

// Sqlite3 数据库连接串
func (a Sqlite3) DSN() string {
	return a.Path
}
