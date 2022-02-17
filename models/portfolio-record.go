package models

import "gorm.io/gorm"

type PortfolioRecord struct {
	gorm.Model
	ProjectID    uint `gorm:"primaryKey"`
	Project      *Project
	UserID       uint `gorm:"primaryKey"`
	User         *User
	Summary      string
	Technologies *[]Technology `gorm:"many2many:portfolio_record_technologies"`
}

type PortfolioRecords []*PortfolioRecord
