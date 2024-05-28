package store

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetDBConnection(dsn string) {
	if dsn == "" {
		log.Panicln("DB options can't be nil")
	} else {
		var err error
		db, err= gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Panicln("Failed to connect to the database", err)
		}
	}
}

func GetDBConnection() *gorm.DB {
	return db
}