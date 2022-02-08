package models

type Role struct {
	ID    uint `gorm:"autoIncrement; primaryKey"`
	Name  string
	Users []*User `gorm:"many2many:user_roles"`
}

type Roles []*Role
