package models

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	Blob      string
	OwnerID   int
	OwnerType string
}
