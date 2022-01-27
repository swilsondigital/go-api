package models

import (
	"encoding/json"
	"math/rand"
	"time"
)

/**
* Setup user struct
 */
type User struct {
	ID              int64     `json:"id" gorm:"primary_key"`
	FirstName       string    `json:"fname"`
	LastName        string    `json:"lname"`
	PreferredName   string    `json:"pname"`
	Email           string    `json:"email"`
	Skillset        string    `json:"skills"`
	YearsExperience int       `json:"experience"`
	MemberSince     time.Time `json:"since"`
}

/**
* get random skill from user skillset
 */
func (u *User) GetRandomSkill() string {
	var skillList []string
	json.Unmarshal([]byte(u.Skillset), &skillList)
	randomSkill := skillList[rand.Intn(len(skillList))]
	return randomSkill
}

/**
* get user full name
 */
func (u *User) GetFullName() string {
	name := u.FirstName + " " + u.LastName
	return name
}
