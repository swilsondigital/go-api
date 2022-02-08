package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName     string  `json:"fname" validate:"required,string"`
	LastName      string  `json:"lname" validate:"required,string"`
	PreferredName string  `json:"pname" validate:"string"`
	Email         string  `json:"email" validate:"required,email" gorm:"unique"`
	Phone         string  `json:"phone" validate:"string"`
	Roles         []*Role `gorm:"many2many:user_roles"`
	Profile       Profile
}

type Users []*User

/**
* Get user's full name
 **/
func (u *User) GetFullName() string {
	return u.FirstName + " " + u.LastName
}
