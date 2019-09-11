package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

var db *gorm.DB


func Init() {
	var err error

	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=maygateway password=99103457 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() *gorm.DB {
	return db
}