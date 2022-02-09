package models

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name             string
	ClientID         uint
	Client           *Client
	Technologies     *[]Technology `gorm:"many2many:project_technologies"`
	Start_Date       time.Time
	Delivery_Date    time.Time
	Private          bool
	PortfolioRecords *[]PortfolioRecord
}

type Projects []*Project
