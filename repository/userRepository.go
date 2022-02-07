package repository

import (
	"math/rand"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

type UserRepository interface {
}

// /**
// * get random skill from user skillset
//  */
// func (u *User) GetRandomSkill() (string, bool) {
// 	var skillList []string
// 	// check if u.Skillset is empty
// 	if u.Skillset == "[]" {
// 		return "", false
// 	}
// 	json.Unmarshal([]byte(u.Skillset), &skillList)
// 	randomSkill := skillList[rand.Intn(len(skillList))]
// 	return randomSkill, true
// }

/**
* get user full name
 */
func (u *User) GetFullName() string {
	name := u.FirstName + " " + u.LastName
	return name
}

/**
* Get All Users - Index
 */
func GetAllUsers(db *gorm.DB) ([]User, error) {
	// setup user query
	var users []User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

/**
* Create New User - Create
 */
func (u *User) CreateNewUser(db *gorm.DB) error {
	// start transaction for creation
	db.Transaction(func(tx *gorm.DB) error {
		// try create
		if err := tx.Create(&u).Error; err != nil {
			return err
		}
		// return nil commits transaction
		return nil
	})
	return nil
}

/**
* Get Single User - Read
 */
func (u *User) GetUser(db *gorm.DB) (User, error) {
	// get model and check db for user
	var user User
	err := db.Where("id = ?", u.ID).First(&user).Error
	return user, err
}

/**
* Get a Random User
 */
func GetRandomUser(db *gorm.DB) (User, error) {
	users, err := GetAllUsers(db)
	if err != nil {
		return User{}, err
	}

	// select random user and return
	randomUser := users[rand.Intn(len(users))]
	return randomUser, nil
}

/**
* Update Single User - Update
 */
func (u *User) UpdateUser(db *gorm.DB, UserInput User) error {
	// start transaction for update
	db.Transaction(func(tx *gorm.DB) error {
		// try create
		if err := tx.Model(&u).Updates(UserInput).Error; err != nil {
			return err
		}
		// return nil commits transaction
		return nil
	})

	return nil
}

/**
* Delete Single User - Delete
 */
func (u *User) DeleteUser(db *gorm.DB) error {

	// start transaction for deletion
	db.Transaction(func(tx *gorm.DB) error {
		// try create
		if err := tx.Delete(&u).Error; err != nil {
			return err
		}
		// return nil commits transaction
		return nil
	})
	return nil
}

/**
* Delete All Users - DeleteAll
 */
func DeleteAllUsers(db *gorm.DB) error {
	users, err := GetAllUsers(db)
	if err != nil {
		return err
	}

	// loop through users models and delete
	for _, user := range users {
		db.Delete(&user)
	}
	return nil
}
