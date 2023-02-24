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
	db, err = gorm.Open(mysql.Open(uri), &gorm.Config{
		CreateBatchSize: 1000,
	})
	if err != nil {
		log.Panic(err.Error())
	}
}

func GetDB() *gorm.DB {
	return db
}
