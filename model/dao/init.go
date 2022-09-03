package dao

import (
	"example.com/http_demo/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var _DB *gorm.DB

func DB() *gorm.DB {
	return _DB
}

func init() {
	_DB = initDB()
}

func initDB() *gorm.DB {
	// In our docker dev environment use
	//db, err := gorm.Open("mysql", "root:superpass@tcp(database:3306)/go_web?charset=utf8&parseTime=True&loc=Local")
	// Out of docker use
	db, err := gorm.Open(config.Database.Type, config.Database.DSN)
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxOpenConns(config.Database.MaxOpenConn)
	db.DB().SetMaxIdleConns(config.Database.MaxIdleConn)
	db.DB().SetConnMaxLifetime(config.Database.MaxLifeTime)
	if err = db.DB().Ping(); err != nil {
		panic(err)
	}
	return db
}
