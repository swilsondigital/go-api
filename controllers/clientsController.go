package controllers

import (
	"goapi/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type clientController struct {
	// place for auth later
	clientRepository repository.ClientRepository
}
type ClientController interface {
	GetAllClients(c *gin.Context)
	GetClientById(c *gin.Context)
	CreateClient(c *gin.Context)
	UpdateClient(c *gin.Context)
	DeleteClient(c *gin.Context)
}

/**
* expected format of json post/put requests
 **/
type ClientInput struct {
	ID              int      `json:"id"`
	FirstName       string   `json:"fname"`
	LastName        string   `json:"lname"`
	PreferredName   string   `json:"pname"`
	Email           string   `json:"email"`
	Phone           string   `json:"phone"`
	Technologies    []string `json:"technologies"`
	YearsExperience int      `json:"experience"`
	MemberSince     string   `json:"since"` // accepts yyyy-mm-dd
}

/**
* Setup New Client Controller
 **/
func NewClientController(cr repository.ClientRepository) ClientController {
	return clientController{clientRepository: cr}
}

/**
* Get All Clients
 **/
func (cc clientController) GetAllClients(c *gin.Context) {
	clients, err := cc.clientRepository.FindAllClients()
	if err != nil {
		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(c.Writer, http.StatusOK, clients)
}

/**
* Get Client by ID
 **/
func (cc clientController) GetClientById(c *gin.Context) {
	id := c.Param("id")
	client, err := cc.clientRepository.FindClientById(id)
	if err != nil {
		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(c.Writer, http.StatusOK, client)
}

/**
* Create New Client
 **/
func (cc clientController) CreateClient(c *gin.Context) {
	// Get POST data
	// var input ClientInput
	// err := json.NewDecoder(c.Request.Body).Decode(&input)
	// // print error if any
	// if err != nil {
	// 	RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// // create Client with repo
	// u, err := cc.clientRepository.CreateClient(Client)
	// if err != nil {
	// 	RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
	// 	return
	// }
	// RespondWithJson(c.Writer, http.StatusOK, u)
}

/**
* Updated Existing Client
 **/
func (cc clientController) UpdateClient(c *gin.Context) {
	// // Get POST data
	// var input ClientInput
	// err := json.NewDecoder(c.Request.Body).Decode(&input)
	// // print error if any
	// if err != nil {
	// 	RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// // get current Client
	// id := c.Param("id")
	// Client, err := cc.clientRepository.FindClientById(id)
	// if err != nil {
	// 	RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// // map input data to Client model
	// newClientModel := models.Client{
	// 	FirstName:     input.FirstName,
	// 	LastName:      input.LastName,
	// 	PreferredName: input.PreferredName,
	// 	Email:         input.Email,
	// 	Phone:         input.Phone,
	// }

	// // load profile if available
	// if input.Technologies != nil || input.YearsExperience != 0 || input.MemberSince != "" {
	// 	var profile models.Profile
	// 	database.DB.Preload("Technologies").Model(&Client).Association("Profile").Find(&profile)

	// 	// parse since time
	// 	since, err := time.Parse("2006-01-02", input.MemberSince)
	// 	if err != nil {
	// 		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
	// 		return
	// 	}

	// 	// assign memberSince and Years Experience
	// 	profile.MemberSince = since
	// 	profile.YearsExperience = input.YearsExperience

	// 	// update associations to technologies
	// 	if input.Technologies != nil {
	// 		var technologies []models.Technology
	// 		for _, v := range input.Technologies {
	// 			var tech models.Technology
	// 			database.DB.Where(models.Technology{Name: v}).FirstOrInit(&tech)
	// 			technologies = append(technologies, tech)
	// 		}

	// 		database.DB.Model(&profile).Association("Technologies").Replace(technologies)
	// 	}

	// 	newClientModel.Profile = profile

	// }
	// // RespondWithJson(c.Writer, http.StatusOK, newClientModel)
	// // update Client
	// updatedClient, err := cc.clientRepository.UpdateClient(Client, newClientModel)
	// if err != nil {
	// 	RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
	// 	return
	// }
	// RespondWithJson(c.Writer, http.StatusOK, updatedClient)
}

/**
* Delete Client
 **/
func (cc clientController) DeleteClient(c *gin.Context) {
	id := c.Param("id")
	err := cc.clientRepository.DeleteClientById(id)
	if err != nil {
		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(c.Writer, http.StatusOK, "Client Deleted Successfully")
}
