package models

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	UserID          uint
	Technologies    []Technology `json:"technology" gorm:"many2many:profile_technologies"`
	YearsExperience int          `json:"experience" validate:"numeric"`
	MemberSince     time.Time    `json:"since" validate:"timestamp,lte"`
	ProfilePhoto    Image        `gorm:"polymorphic:Owner;"`
}
