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

type portfolioRecordController struct {
	// place for auth later
	recordRepository repository.PortfolioRecordRepository
}
type PortfolioRecordController interface {
}

/**
* Setup New Project Controller
 **/
func NewPortfolioRecordController(rr repository.PortfolioRecordRepository) PortfolioRecordController {
	return portfolioRecordController{recordRepository: rr}
}

/**
* Get All Project Records
 **/
func (rc portfolioRecordController) GetRecordsByProject(c *gin.Context) {
	id := c.Param("id")
	records, err := rc.recordRepository.FindRecordsByProjectId(id)
	if err != nil {
		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(c.Writer, http.StatusOK, records)
}

/**
* Get Single Record By
 **/
func (rc portfolioRecordController) GetRecordById(c *gin.Context) {
	id := c.Param("id")
	record, err := rc.recordRepository.FindRecordById(id)
	if err != nil {
		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(c.Writer, http.StatusOK, record)
}

/**
* Create New Portfolio Record
 **/
func (rc portfolioRecordController) CreateRecord(c *gin.Context) {
	// Get POST data
	var input ProjectInput
	err := json.NewDecoder(c.Request.Body).Decode(&input)
	// print error if any
	if err != nil {
		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
	// get ClientID from url
	clientID := c.Param("id")

	// check for dates and parse
	startDate, _ := time.Parse("2006-01-02", input.Start_Date)
	deliveryDate, _ := time.Parse("2006-01-02", input.Delivery_Date)

	// map data to new project type
	project := models.Project{
		Name:          input.Name,
		Start_Date:    startDate,
		Delivery_Date: deliveryDate,
		Private:       input.Private,
	}

	// add project technologies
	if input.Technologies != nil {
		var technologies []models.Technology
		for _, v := range input.Technologies {

			// TODO: Add check for lowercase/uppercase/caps
			var tech models.Technology
			database.DB.Where(models.Technology{Name: v}).FirstOrInit(&tech)
			technologies = append(technologies, tech)
		}

		project.Technologies = &technologies
	}

	// get client model from db
	var client models.Client
	err = database.DB.Where("id = ?", clientID).First(&client).Error
	if err != nil {
		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}

	// associate project to client
	project.Client = &client

	// update project to private if client is set as private
	if client.Private && !project.Private {
		project.Private = true
	}

	// create user with repo
	p, err := rc.recordRepository.CreateProject(project)
	if err != nil {
		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(c.Writer, http.StatusOK, p)
}

/**
* Update Project
 **/
func (rc portfolioRecordController) UpdateProject(c *gin.Context) {
	// Get POST data
	var input ProjectInput
	err := json.NewDecoder(c.Request.Body).Decode(&input)
	// print error if any
	if err != nil {
		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}

	// get current project
	id := c.Param("id")
	project, err := rc.recordRepository.FindProjectById(id)
	if err != nil {
		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}

	// check for dates and parse
	startDate, _ := time.Parse("2006-01-02", input.Start_Date)
	deliveryDate, _ := time.Parse("2006-01-02", input.Delivery_Date)

	// map data to new project type
	newProjectModel := models.Project{
		Name:          input.Name,
		Start_Date:    startDate,
		Delivery_Date: deliveryDate,
		Private:       input.Private,
	}

	// add project technologies
	if input.Technologies != nil {
		var technologies []models.Technology
		for _, v := range input.Technologies {

			// TODO: Add check for lowercase/uppercase/caps
			var tech models.Technology
			database.DB.Where(models.Technology{Name: v}).FirstOrInit(&tech)
			technologies = append(technologies, tech)
		}

		newProjectModel.Technologies = &technologies
	}

	// update project to private if client is set as private
	if project.Client.Private && !newProjectModel.Private {
		newProjectModel.Private = true
	}

	// create user with repo
	p, err := rc.recordRepository.UpdateProject(project, newProjectModel)
	if err != nil {
		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(c.Writer, http.StatusOK, p)
}

/**
* Delete Project
 **/
func (rc portfolioRecordController) DeleteProject(c *gin.Context) {
	id := c.Param("id")
	err := rc.recordRepository.DeleteProjectById(id)
	if err != nil {
		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(c.Writer, http.StatusOK, "Project Deleted Successfully")
}
