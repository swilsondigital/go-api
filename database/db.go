package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

/**
* Initialize DB table
 */
func InitDB(postgresURL string) {
	// setup connection to db
	db, err := gorm.Open(postgres.Open(postgresURL), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
}

/**
* Get Connection to DB
 */
func GetDB() *gorm.DB {
	return DB
}
