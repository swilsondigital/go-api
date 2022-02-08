package models

import "gorm.io/gorm"

type PortfolioRecord struct {
	gorm.Model
	ClientID     uint `gorm:"primaryKey"`
	ProjectID    uint `gorm:"primaryKey"`
	Summary      string
	Technologies []Technology `gorm:"many2many:portfolio_record_technologies"`
}

type PortfolioRecords []*PortfolioRecord
