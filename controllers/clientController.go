package controllers

import (
	"encoding/json"
	"fmt"
	"goapi/database"
	"goapi/models"
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
	Name    string
	Phone   string
	Private bool
	Logo    string
	Contact map[string]string
	Address map[string]string
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
	var input ClientInput
	err := json.NewDecoder(c.Request.Body).Decode(&input)
	// print error if any
	if err != nil {
		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}

	// map data to new client type
	client := models.Client{
		Name:    input.Name,
		Phone:   input.Phone,
		Private: input.Private,
	}

	fmt.Println(len(input.Address))

	// check for client address
	if len(input.Address) != 0 {
		address := models.Address{
			Address_1:      input.Address["Address_1"],
			Address_2:      input.Address["Address_2"],
			City:           input.Address["City"],
			State_Province: input.Address["State_Province"],
			Postal_Code:    input.Address["Postal_Code"],
			Country:        input.Address["Country"],
		}
		client.Address = &address
	}

	// check for client contacts
	if len(input.Contact) != 0 {
		user := models.User{
			FirstName:     input.Contact["FirstName"],
			LastName:      input.Contact["LastName"],
			PreferredName: input.Contact["PreferredName"],
			Email:         input.Contact["Email"],
			Phone:         input.Contact["Phone"],
		}
		database.DB.Where(user).FirstOrInit(&user)
		client.Contact = &user
	}

	// check for client logo
	if input.Logo != "" {
		logo := models.Image{
			Blob: input.Logo,
		}
		client.Logo = &logo
	}

	// create Client with repo
	u, err := cc.clientRepository.CreateClient(client)
	if err != nil {
		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(c.Writer, http.StatusOK, u)
}

/**
* Updated Existing Client
 **/
func (cc clientController) UpdateClient(c *gin.Context) {
	// Get POST data
	var input ClientInput
	err := json.NewDecoder(c.Request.Body).Decode(&input)
	// print error if any
	if err != nil {
		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}

	// get current client
	id := c.Param("id")
	client, err := cc.clientRepository.FindClientById(id)
	if err != nil {
		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}

	newClientModel := models.Client{
		Name:    input.Name,
		Phone:   input.Phone,
		Private: input.Private,
	}

	// check for client address
	if len(input.Address) != 0 {
		address := models.Address{
			Address_1:      input.Address["Address_1"],
			Address_2:      input.Address["Address_2"],
			City:           input.Address["City"],
			State_Province: input.Address["State_Province"],
			Postal_Code:    input.Address["Postal_Code"],
			Country:        input.Address["Country"],
		}
		newClientModel.Address = &address
	}

	// check for client contacts
	if len(input.Contact) != 0 {
		user := models.User{
			FirstName:     input.Contact["FirstName"],
			LastName:      input.Contact["LastName"],
			PreferredName: input.Contact["PreferredName"],
			Email:         input.Contact["Email"],
			Phone:         input.Contact["Phone"],
		}
		database.DB.Where(user).FirstOrInit(&user)
		newClientModel.Contact = &user
	}

	// check for client logo
	if input.Logo != "" {
		logo := models.Image{
			Blob: input.Logo,
		}
		newClientModel.Logo = &logo
	}

	updatedClient, err := cc.clientRepository.UpdateClient(client, newClientModel)
	if err != nil {
		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(c.Writer, http.StatusOK, updatedClient)

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
