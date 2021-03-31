package database

import (
	"log"

	"github.com/olivere/elastic"

	"github.com/deposit-services/api"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	elIndex = "deposit"
)

var db *gorm.DB
var elasticConn *elastic.Client

func init() {
	conn, err := gorm.Open("mysql", "root@(localhost:3307)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err)
	}
	db = conn

	client, elasticErr := elastic.NewClient(elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))
	if elasticErr != nil {
		log.Fatalln(elasticErr)
	}
	elasticConn = client

	Migrate()
}

func Migrate() {
	errDp := db.Debug().AutoMigrate(&api.Deposit{})
	if errDp != nil {
		log.Println(errDp)
	}
	errEv := db.Debug().AutoMigrate(&api.EventStore{})
	if errEv != nil {
		log.Println(errEv)
	}
}
