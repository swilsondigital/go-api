package models

import "gorm.io/gorm"

type PortfolioRecord struct {
	gorm.Model
	ProjectID    uint `gorm:"primaryKey"`
	UserID       uint `gorm:"primaryKey"`
	Summary      string
	Technologies *[]Technology `gorm:"many2many:portfolio_record_technologies"`
}

type PortfolioRecords []*PortfolioRecord
