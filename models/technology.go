package models

import "gorm.io/gorm"

type Technology struct {
	gorm.Model
	Name string
}
