package mock

import (
	"errors"
	"fmt"
	"goapi/models"
	"math/rand"
)

type User struct {
	models.User
}

type MockDB []User

/**
* Get All Users - Index
 */
func MockModelGetAllUsers(db MockDB) ([]User, error) {
	return db, nil
}

/**
* Create New User - Create
 */
func (u User) MockModelCreateNewUser(db MockDB) error {
	db = append(db, u)
	return nil
}

/**
* Get Single User - Read
 */
func (u User) MockModelGetUser(db MockDB) (User, error) {
	var user User
	// setup error incase user doesnt exist
	err := errors.New(fmt.Sprintf("User with ID %d could not be found", u.ID))
	for _, v := range db {
		if v.ID == u.ID {
			user = v
			err = nil
			break
		}
	}

	return user, err
}

/**
* Get a Random User
 */
func MockModelGetRandomUser(db MockDB) (User, error) {
	users := db
	// select random user and return
	randomUser := users[rand.Intn(len(users))]
	return randomUser, nil
}

/**
* Update Single User - Update
 */
func (u User) MockModelUpdateUser(db MockDB, UserInput User) error {
	user, err := u.MockModelGetUser(db)
	if err != nil {
		return err
	}

	for key, v := range db {
		if user.ID == v.ID {
			db[key] = UserInput
			break
		}
	}
	return nil
}

/**
* Delete Single User - Delete
 */
func (u User) MockModelDeleteUser(db MockDB) error {

	user, err := u.MockModelGetUser(db)
	if err != nil {
		return err
	}

	var index int

	for key, v := range db {
		if user.ID == v.ID {
			index = key
			break
		}
	}
	// replace with last element
	db[index] = db[len(db)-1]
	// remove last element
	db = db[:len(db)-1]

	return nil
}

/**
* Delete All Users - DeleteAll
 */
func MockModelDeleteAllUsers(db MockDB) error {
	db = []User{}
	return nil
}
