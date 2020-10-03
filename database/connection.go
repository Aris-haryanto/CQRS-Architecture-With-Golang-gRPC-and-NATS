package database

import (
	"log"

	"github.com/deposit-services/api"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	conn, err := gorm.Open("mysql", "root@(localhost)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err)
	}

	db = conn

	Migrate()
}

func Migrate() {
	err := db.Debug().AutoMigrate(&api.Deposit{})
	if err != nil {
		log.Println(err)
	}
}
