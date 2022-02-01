package database

import (
	"goapi/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

/**
* Initialize DB table
 */
func InitDB(postgresURL string) *gorm.DB {
	// setup connection to db
	db, err := gorm.Open(postgres.Open(postgresURL), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate([]models.User{})

	DB = db
	return DB
}

/**
* Get Connection to DB
 */
func GetDB() *gorm.DB {
	return DB
}
