package controllers

import (
	"encoding/json"
	"fmt"
	"goapi/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type UserController struct {
	Router *mux.Router
	DB     *gorm.DB
}

type UserInput struct {
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
func (c *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetAllUsers(c.DB)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, users)
}

/**
* Create New User - Create
 */
func (c *UserController) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	var input UserInput
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

	if err := user.CreateNewUser(c.DB); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, user)
}

/**
* Get Single User - Read
 */
func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 0)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}

	user := models.User{}
	user.ID = uint(id)
	resp, err := user.GetUser(c.DB)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, resp)
}

/**
* Get single user and show formatted data
 */
func (c *UserController) IntroduceUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 0)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}

	user := models.User{}
	user.ID = uint(id)
	u, err := user.GetUser(c.DB)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	greeting := "Hello! My name is " + u.GetFullName() + "."
	var name string
	// add preferred name to greeting
	if u.PreferredName != "" {
		name = u.PreferredName
	} else {
		name = u.FirstName
	}
	greeting += " You can call me " + name + "."
	// check for random skill
	if randSkill, ok := u.GetRandomSkill(); ok {
		greeting += " I have experience with " + randSkill + "."
	}
	RespondWithJson(w, http.StatusOK, greeting)
}

/**
* Get a Random User
 */
func (c *UserController) GetRandomUser(w http.ResponseWriter, r *http.Request) {
	user, err := models.GetRandomUser(c.DB)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Could not get random user")
		return
	}
	// select random user and return
	RespondWithJson(w, http.StatusOK, user)
}

/**
* Update Single User - Update
 */
func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 0)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}

	user := models.User{}
	user.ID = uint(id)
	u, err := user.GetUser(c.DB)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var input UserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		fmt.Println(err)
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

	if err := u.UpdateUser(c.DB, userInput); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, user)
}

/**
* Delete Single User - Delete
 */
func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 0)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}

	user := models.User{}
	user.ID = uint(id)
	u, err := user.GetUser(c.DB)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := u.DeleteUser(c.DB); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, u)
}

/**
* Delete All Users - DeleteAll
 */
func (c *UserController) DeleteAllUsers(w http.ResponseWriter, r *http.Request) {
	if err := models.DeleteAllUsers(c.DB); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, "All Users Deleted")
}

/**
* Initialize Routes
 */
func (c *UserController) InitializeUserRoutes() {
	// get random user
	c.Router.HandleFunc("/random-user", c.GetRandomUser).Methods("GET")
	// user subrouter paths
	userRouter := c.Router.PathPrefix("/users").Subrouter()
	// create new user
	userRouter.HandleFunc("/", c.CreateNewUser).Methods("POST")
	// get all users
	userRouter.HandleFunc("/", c.GetAllUsers).Methods("GET")
	// get single user
	userRouter.HandleFunc("/{id}", c.GetUser).Methods("GET")
	// update single user
	userRouter.HandleFunc("/{id}", c.UpdateUser).Methods("PUT")
	// delete single user
	userRouter.HandleFunc("/{id}", c.DeleteUser).Methods("DELETE")
	// get single user
	userRouter.HandleFunc("/{id}/hello", c.IntroduceUser).Methods("GET")
	// delete all users
	userRouter.HandleFunc("/", c.DeleteAllUsers).Methods("DELETE")
}
