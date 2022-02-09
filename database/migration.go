package database

import "goapi/models"

func Automigrate() {
	err := DB.AutoMigrate(
		&models.Technology{},
		&models.Image{},
		&models.User{},
		&models.Client{},
		&models.Address{},
		&models.PortfolioRecord{},
		&models.Profile{},
		&models.Project{},
	)
	if err != nil {
		panic(err)
	}
}
