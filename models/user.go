package models

import (
	"gorm.io/gorm"
)

/**
* Setup user struct
 */
type User struct {
	gorm.Model
	FirstName     string `json:"fname" validate:"required,string"`
	LastName      string `json:"lname" validate:"required,string"`
	PreferredName string `json:"pname" validate:"string"`
	Email         string `json:"email" validate:"required,email" gorm:"unique"`
	Phone         string `json:"phone" validate:"string"`
	RoleID        int
	Role          Role `gorm:"foreignKey:RoleID"`
}
