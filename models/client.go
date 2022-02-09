package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	Name      string
	Logo      *Image `gorm:"polymorphic:Owner;" json:",omitempty"`
	Phone     string
	ContactID uint
	Contact   *User `gorm:"foreignKey:ContactID" json:",omitempty"`
	Private   bool
	Projects  *[]Project `json:",omitempty"`
	Address   *Address   `json:",omitempty"`
}

type Clients []*Client
