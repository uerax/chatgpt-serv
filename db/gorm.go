package db

import (
	"github.com/uerax/goconf"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var log *zap.Logger

func Init() {
	connect()
}

func connect() {
	uri, err := goconf.VarString("database", "mysql", "uri")
	if err != nil {
		log.Panic(err.Error())
	}
	db, err = gorm.Open(mysql.Open(uri), &gorm.Config{})
	if err != nil {
		log.Panic(err.Error())
	}
}

func Ping() bool {
	d, err := db.DB()
	if err != nil {
		log.Error(err.Error())
	}
	err = d.Ping()
	if err != nil {
		log.Error(err.Error())
		return false
	}

	return true
}