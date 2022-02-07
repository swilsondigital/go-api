package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	Name      string
	Logo      Image `gorm:"polymorphic:Owner;"`
	Phone     string
	ContactID int
	Contact   User `gorm:"foreignKey:ContactID"`
	Private   bool
}
