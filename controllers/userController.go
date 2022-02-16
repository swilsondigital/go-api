package controllers

import (
	"encoding/json"
	"goapi/database"
	"goapi/models"
	"goapi/repository"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type userController struct {
	// place for auth later
	userRepository repository.UserRepository
}
type UserController interface {
	GetAllUsers(c *gin.Context)
	GetUserById(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

/**
* expected format of json post/put requests
 **/
type UserInput struct {
	FirstName       string
	LastName        string
	PreferredName   string
	Email           string
	Phone           string
	Technologies    []string
	YearsExperience int
	MemberSince     string // accepts yyyy-mm-dd
	ProfilePhoto    string
}

/**
* Setup New User Controller
 **/
func NewUserController(ur repository.UserRepository) UserController {
	return userController{userRepository: ur}
}

/**
* Get All Users
 **/
func (uc userController) GetAllUsers(c *gin.Context) {
	users, err := uc.userRepository.FindAllUsers()
	if err != nil {
		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(c.Writer, http.StatusOK, users)
}

/**
* Get User By ID
 **/
func (uc userController) GetUserById(c *gin.Context) {
	id := c.Param("id")
	user, err := uc.userRepository.FindUserById(id)
	if err != nil {
		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(c.Writer, http.StatusOK, user)
}

/**
* Create New User
 **/
func (uc userController) CreateUser(c *gin.Context) {
	// Get POST data
	var input UserInput
	err := json.NewDecoder(c.Request.Body).Decode(&input)
	// print error if any
	if err != nil {
		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}

	// map input data to user model
	user := models.User{
		FirstName:     input.FirstName,
		LastName:      input.LastName,
		PreferredName: input.PreferredName,
		Email:         input.Email,
		Phone:         input.Phone,
	}

	if input.Technologies != nil || input.YearsExperience != 0 || input.MemberSince != "" || input.ProfilePhoto != "" {

		// init profile model
		profile := models.Profile{
			YearsExperience: input.YearsExperience,
		}

		// parse member since
		since, err := time.Parse("2006-01-02", input.MemberSince)
		if err != nil {
			RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
			return
		} else {
			profile.MemberSince = &since
		}

		// check for technologies
		if input.Technologies != nil {
			var technologies []models.Technology
			for _, v := range input.Technologies {

				// TODO: Add check for lowercase/uppercase/caps
				var tech models.Technology
				database.DB.Where(models.Technology{Name: v}).FirstOrInit(&tech)
				technologies = append(technologies, tech)
			}
			profile.Technologies = &technologies
		}

		// check for profile photo
		if input.ProfilePhoto != "" {
			photo := models.Image{
				Blob: input.ProfilePhoto,
			}
			profile.ProfilePhoto = &photo
		}

		// TODO: check if profile should be created (a client contact)
		user.Profile = &profile
	}

	// TODO: assign roles

	// create user with repo
	u, err := uc.userRepository.CreateUser(user)
	if err != nil {
		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(c.Writer, http.StatusOK, u)
}

/**
* Updated Existing User
 **/
func (uc userController) UpdateUser(c *gin.Context) {
	// Get POST data
	var input UserInput
	err := json.NewDecoder(c.Request.Body).Decode(&input)
	// print error if any
	if err != nil {
		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}

	// get current User
	id := c.Param("id")
	user, err := uc.userRepository.FindUserById(id)
	if err != nil {
		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}

	// map input data to user model
	newUserModel := models.User{
		FirstName:     input.FirstName,
		LastName:      input.LastName,
		PreferredName: input.PreferredName,
		Email:         input.Email,
		Phone:         input.Phone,
	}

	// load profile if available
	if input.Technologies != nil || input.YearsExperience != 0 || input.MemberSince != "" || input.ProfilePhoto != "" {
		var profile models.Profile
		database.DB.Preload("Technologies").Model(&user).Association("Profile").Find(&profile)

		// add years experience
		if input.YearsExperience != 0 {
			profile.YearsExperience = input.YearsExperience
		}

		// check member since
		if input.MemberSince != "" {
			// parse since time
			since, err := time.Parse("2006-01-02", input.MemberSince)
			if err != nil {
				RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
				return
			}
			profile.MemberSince = &since
		}

		// update associations to technologies
		if input.Technologies != nil {
			var technologies []models.Technology
			for _, v := range input.Technologies {
				// TODO: Add check for lowercase/uppercase/caps
				var tech models.Technology
				database.DB.Where(models.Technology{Name: v}).FirstOrInit(&tech)
				technologies = append(technologies, tech)
			}

			database.DB.Model(&profile).Association("Technologies").Replace(&technologies)
		}

		// check for profile photo
		if input.ProfilePhoto != "" {
			photo := models.Image{
				Blob: input.ProfilePhoto,
			}
			database.DB.Model(&profile).Association("ProfilePhoto").Replace(&photo)
		}

		newUserModel.Profile = &profile

	}
	// update user
	updatedUser, err := uc.userRepository.UpdateUser(user, newUserModel)
	if err != nil {
		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(c.Writer, http.StatusOK, updatedUser)
}

/**
* Delete User
 **/
func (uc userController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := uc.userRepository.DeleteUserById(id)
	if err != nil {
		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(c.Writer, http.StatusOK, "User Deleted Successfully")
}
