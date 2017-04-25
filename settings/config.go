package settings

import (
	"fmt"
)

const adapterNameMysql string = "mysql"
const adapterNameMysql2 string = "mysql2"

type Config struct {
	DB        *DBConfig        `mapstructure:"database"`
	Batch     *BatchConfig     `mapstructure:"batch"`
	AppServer *AppServerConfig `mapstructure:"appServer"`
	ApiServer *ApiServerConfig `mapstructure:"apiServer"`
}

type AppServerConfig struct {
	LogFile string
	Port    string
	PidFile string
}

type ApiServerConfig struct {
	LogFile string
	Port    string
	PidFile string
}

type BatchConfig struct {
	LogFile string
}

type DBConfig struct {
	Adapter  string
	Charset  string
	Encoding string
	Database string
	Username string
	Password string
	Host     string
	Port     int
	Pool     int
	LogFile  string
}

// AdapterName get db adapter name
func (db *DBConfig) AdapterName() string {
	switch db.Adapter {
	case adapterNameMysql, adapterNameMysql2:
		return adapterNameMysql
	}
	return ""
}

// DSN get db dns
func (db *DBConfig) DSN() string {
	switch db.Adapter {
	case adapterNameMysql, adapterNameMysql2:
		return fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
			db.Username,
			db.Password,
			db.Host,
			db.Port,
			db.Database,
			db.Charset,
		)
	}
	return ""
}
