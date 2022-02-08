package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	Name    string
	Logo    Image `gorm:"polymorphic:Owner;"`
	Phone   string
	Contact uint
	User    User `gorm:"foreignKey:Contact"`
	Private bool
}

type Clients []*Client
