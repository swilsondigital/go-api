package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Address_1      string
	Address_2      string
	City           string
	State_Province string
	Postal_Code    string
	Country        string
	ClientID       uint
}
