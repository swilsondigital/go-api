package models

import "gorm.io/gorm"

type PortfolioRecord struct {
	gorm.Model
	ClientID     int `gorm:"primaryKey"`
	ProjectID    int `gorm:"primaryKey"`
	Summary      string
	Technologies []Technology `gorm:"many2many:portfolio_record_technologies"`
}
