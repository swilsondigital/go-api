package models

type Role struct {
	ID   int `gorm:"autoIncrement; primaryKey"`
	Name string
}
