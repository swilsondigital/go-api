package users

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	db "goapi/src"

	"github.com/gorilla/mux"
)

/**
* Setup user struct
 */
type User struct {
	// this should be auto incremented so probably based on the db
	ID              int64     `json:"id" field:"id, identity" validate:"unique=ID"`
	FirstName       string    `json:"fname"`
	LastName        string    `json:"lname"`
	PreferredName   string    `json:"pname"`
	Email           string    `json:"email" validate:"required,email"`           // required email
	Skillset        []string  `json:"skills" validate:"gt=0,dive,dive,required"` // greater than 0 [] & string is required
	YearsExperience int       `json:"experience" validate:"gt=0,lt=130"`         // between 0-130
	MemberSince     time.Time `json:"since" validate:"lte"`                      // less than today
}

/**
* get random skill from user skillset
 */
func (u User) getRandomSkill() string {
	randomSkill := u.Skillset[rand.Intn(len(u.Skillset))]
	return randomSkill
}

/**
* get user full name
 */
func (u User) getFullName() string {
	name := u.FirstName + " " + u.LastName
	return name
}

/**
* Init users
 */
var Users []User

/**
* Get All Users - Index
 */
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Returning all users")
	fmt.Println(len(Users))
	json.NewEncoder(w).Encode(&Users)
}

/**
* Create New User - Create
 */
func CreateNewUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	// print error if any
	if err != nil {
		fmt.Println(err)
		return
	}

	var keys []string
	var values []interface{}

	// setup user reflect
	u := reflect.ValueOf(&user).Elem()
	typeOfUser := u.Type()

	// iterate over user reflection
	for i := 0; i < u.NumField(); i++ {
		f := u.Field(i)
		keys = append(keys, typeOfUser.Field(i).Name)
		values = append(values, f.Interface())
	}

	keysToString := strings.Join(keys[:], ",")
	valuesToString := strings.Trim
	createUserQuery := "INSERT INTO users (%k) VALUES (%v);"

	res := db.RunQuery(fmt.Sprintf(createUserQuery, strings.Join(keys[:], ",")))

	// add user to Users slice
	Users = append(Users, user)
	// return the new Users slice
	json.NewEncoder(w).Encode(&Users)
}

/**
* Get Single User - Read
 */
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("Returning user with ID:", id)

	for _, user := range Users {
		// if we find the user with the same id
		idInt, _ := strconv.ParseInt(id, 10, 64)
		if user.ID == idInt {
			json.NewEncoder(w).Encode(&user)
			return
		}
	}
	// return message
	json.NewEncoder(w).Encode("User with id: " + id + " could not be found")
}

/**
* Get single user and show formatted data
 */
func IntroduceUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("Returning user with ID:", id)

	for _, user := range Users {
		// if we find the user with the same id
		idInt, _ := strconv.ParseInt(id, 10, 64)
		if user.ID == idInt {
			greeting := "Hello! My name is " + user.getFullName() + ". You can call me " + user.PreferredName + ". I have programming experience with " + user.getRandomSkill() + "."
			json.NewEncoder(w).Encode(&greeting)
			return
		}
	}

	// return message
	json.NewEncoder(w).Encode("User with id: " + id + " could not be found")
}

/**
* Get a Random User
 */
func GetRandomUser(w http.ResponseWriter, r *http.Request) {
	if len(Users) == 0 {
		// return message
		json.NewEncoder(w).Encode("No users were found")
	}
	// select random user and return
	randomUser := Users[rand.Intn(len(Users))]
	json.NewEncoder(w).Encode(&randomUser)
}

/**
* Update Single User - Update
 */
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("Retrieving user with ID:", id)

	var userKeyToUpdate int
	var updatedUser User
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	// print error if any
	if err != nil {
		fmt.Println(err)
		return
	}

	// loop through to find the key for the user in the slice
	for key, user := range Users {
		// if we find the user with the same id
		idInt, _ := strconv.ParseInt(id, 10, 64)
		if user.ID == idInt {
			userKeyToUpdate = key

			// add user to Users slice
			Users[userKeyToUpdate] = updatedUser
			// return an empty user object
			json.NewEncoder(w).Encode(&updatedUser)
			return
		}
	}
	// no user to update
	json.NewEncoder(w).Encode("User with id: " + id + " could not be found")
}

/**
* Delete Single User - Delete
 */
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("Deleting user with ID:", id)

	var userKeyToDelete int

	// loop through to find the key for the user in the slice
	for key, user := range Users {
		// if we find the user with the same id
		idInt, _ := strconv.ParseInt(id, 10, 64)
		if user.ID == idInt {
			userKeyToDelete = key

			// update Users list
			Users = append(Users[:userKeyToDelete], Users[userKeyToDelete+1:]...)
			json.NewEncoder(w).Encode(&Users)
			return
		}
	}
	// no user to delete
	json.NewEncoder(w).Encode("User with id: " + id + " could not be found")
}

/**
* Delete All Users - DeleteAll
 */
func DeleteAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting All Users")
	Users = nil
	json.NewEncoder(w).Encode("All Users Deleted")
}
