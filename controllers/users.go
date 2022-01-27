package userController

import (
	"encoding/json"
	"fmt"
	"goapi/models"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type CreateUserInput struct {
	FirstName       string    `json:"fname"`
	LastName        string    `json:"lname"`
	PreferredName   string    `json:"pname"`
	Email           string    `json:"email"`
	Skillset        []string  `json:"skills"`
	YearsExperience int       `json:"experience"`
	MemberSince     time.Time `json:"since"`
}

type UpdateUserInput struct {
	FirstName       string    `json:"fname"`
	LastName        string    `json:"lname"`
	PreferredName   string    `json:"pname"`
	Email           string    `json:"email"`
	Skillset        []string  `json:"skills"`
	YearsExperience int       `json:"experience"`
	MemberSince     time.Time `json:"since"`
}

/**
* Get All Users - Index
 */
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// setup user query
	var users []models.User
	models.DB.Find(&users)
	// return message
	fmt.Println("Returning all users")
	fmt.Println(len(users))
	json.NewEncoder(w).Encode(&users)
}

/**
* Create New User - Create
 */
func CreateNewUser(w http.ResponseWriter, r *http.Request) {
	var input CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&input)
	// print error if any
	if err != nil {
		fmt.Println(err)
		return
	}

	// convert skillset data to marshalled json
	skills, _ := json.Marshal(input.Skillset)

	// map input data to user model
	user := models.User{
		FirstName:       input.FirstName,
		LastName:        input.LastName,
		PreferredName:   input.PreferredName,
		Email:           input.Email,
		Skillset:        string(skills),
		YearsExperience: input.YearsExperience,
		MemberSince:     input.MemberSince,
	}

	// start transaction for creation
	models.DB.Transaction(func(tx *gorm.DB) error {
		// try create
		if err := tx.Create(&user).Error; err != nil {
			return err
		}
		// return nil commits transaction
		return nil
	})
	json.NewEncoder(w).Encode(user)
}

/**
* Get Single User - Read
 */
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("Returning user with ID:", id)

	// get model and check db for user
	var user models.User
	if err := models.DB.Where("id = ?", id).First(&user).Error; err != nil {
		// return message about no user found
		json.NewEncoder(w).Encode("User with id: " + id + " could not be found")
		return
	}
	// return user
	json.NewEncoder(w).Encode(&user)
}

/**
* Get single user and show formatted data
 */
func IntroduceUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("Returning user with ID:", id)

	// get model and check db for user
	var user models.User
	if err := models.DB.Where("id = ?", id).First(&user).Error; err != nil {
		// return message about no user found
		json.NewEncoder(w).Encode("User with id: " + id + " could not be found")
		return
	}
	// return user
	greeting := "Hello! My name is " + user.GetFullName() + ". You can call me " + user.PreferredName + ". I have programming experience with " + user.GetRandomSkill() + "."
	json.NewEncoder(w).Encode(&greeting)
}

/**
* Get a Random User
 */
func GetRandomUser(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	models.DB.Find(&users)
	// check if users found
	if len(users) == 0 {
		// return message
		json.NewEncoder(w).Encode("No users were found")
		return
	}

	// select random user and return
	randomUser := users[rand.Intn(len(users))]
	json.NewEncoder(w).Encode(&randomUser)
}

/**
* Update Single User - Update
 */
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("Retrieving user with ID:", id)

	var input UpdateUserInput
	err := json.NewDecoder(r.Body).Decode(&input)
	// print error if any
	if err != nil {
		fmt.Println(err)
		return
	}

	// get model and check db for user
	var user models.User
	if err := models.DB.Where("id = ?", id).First(&user).Error; err != nil {
		// return message about no user found
		json.NewEncoder(w).Encode("User with id: " + id + " could not be found")
		return
	}

	// convert skillset data to marshalled json
	skills, _ := json.Marshal(input.Skillset)

	// map input data to user model
	userInput := models.User{
		FirstName:       input.FirstName,
		LastName:        input.LastName,
		PreferredName:   input.PreferredName,
		Email:           input.Email,
		Skillset:        string(skills),
		YearsExperience: input.YearsExperience,
		MemberSince:     input.MemberSince,
	}

	// start transaction for update
	models.DB.Transaction(func(tx *gorm.DB) error {
		// try create
		if err := tx.Model(&user).Updates(userInput).Error; err != nil {
			return err
		}
		// return nil commits transaction
		return nil
	})

	// return user
	json.NewEncoder(w).Encode(&user)
}

/**
* Delete Single User - Delete
 */
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("Deleting user with ID:", id)

	// get model and check db for user
	var user models.User
	if err := models.DB.Where("id = ?", id).First(&user).Error; err != nil {
		// return message about no user found
		json.NewEncoder(w).Encode("User with id: " + id + " could not be found")
		return
	}

	// start transaction for deletion
	models.DB.Transaction(func(tx *gorm.DB) error {
		// try create
		if err := tx.Delete(&user).Error; err != nil {
			json.NewEncoder(w).Encode("User with id: " + id + " could not be found")
			return err
		}
		// return nil commits transaction
		return nil
	})

	json.NewEncoder(w).Encode("User with id: " + id + " was deleted")
}

/**
* Delete All Users - DeleteAll
 */
func DeleteAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting All Users")

	// get users models
	var users []models.User
	models.DB.Find(&users)

	// loop through users models and delete
	for _, user := range users {
		models.DB.Delete(&user)
	}

	json.NewEncoder(w).Encode("All Users Deleted")
}
