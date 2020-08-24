package db

import (
	"danmu/services/config"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var (
	DB *xorm.Engine
)

func init() {
	connStr := config.Global.MySQL
	conn, err := xorm.NewEngine("mysql", connStr)
	if err != nil {
		panic(err)
	}

	DB = conn
}
