package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
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
	db, err := gorm.Open("mysql", "root:superpass@tcp(localhost:30306)/go_web?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxOpenConns(100)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetConnMaxLifetime(time.Second * 300)
	if err = db.DB().Ping(); err != nil {
		panic(err)
	}
	return db
}
