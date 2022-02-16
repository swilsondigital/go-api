package models

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	UserID          uint
	Technologies    *[]Technology `gorm:"many2many:profile_technologies"`
	YearsExperience int           `validate:"numeric"`
	MemberSince     *time.Time    `validate:"timestamp,lte"`
	ProfilePhoto    *Image        `gorm:"polymorphic:Owner;"`
}
