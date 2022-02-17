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
	Start_Date       *time.Time    `gorm:"default:null"`
	Delivery_Date    *time.Time    `gorm:"default:null"`
	Private          bool
	PortfolioRecords *[]PortfolioRecord `json:",omitempty"`
}

type Projects []*Project
