package database

import "goapi/models"

func Automigrate() {
	err := DB.AutoMigrate(
		&models.Address{},
		&models.Client{},
		&models.Image{},
		&models.PortfolioRecord{},
		&models.Profile{},
		&models.Project{},
		&models.Technology{},
		&models.User{},
	)
	if err != nil {
		panic(err)
	}
}
