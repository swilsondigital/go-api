package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName        string             `validate:"required,string"`
	LastName         string             `validate:"required,string"`
	PreferredName    string             `validate:"string"`
	Email            string             `validate:"required,email" gorm:"unique"`
	Phone            string             `validate:"string"`
	Roles            []*Role            `gorm:"many2many:user_roles" json:",omitempty"`
	Profile          *Profile           `json:",omitempty"`
	PortfolioRecords *[]PortfolioRecord `json:",omitempty"`
}

type Users []*User

/**
* Get user's full name
 **/
func (u *User) GetFullName() string {
	return u.FirstName + " " + u.LastName
}
