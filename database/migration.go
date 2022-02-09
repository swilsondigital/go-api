package database

import "goapi/models"

func Automigrate() {
	err := DB.AutoMigrate(
		&models.Technology{},
		&models.Image{},
		&models.User{},
		&models.Role{},
		&models.PublicPassword{},
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
