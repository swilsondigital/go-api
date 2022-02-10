package controllers

import (
	"encoding/json"
	"goapi/database"
	"goapi/models"
	"goapi/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type portfolioRecordController struct {
	// place for auth later
	recordRepository repository.PortfolioRecordRepository
}
type PortfolioRecordController interface {
	GetRecordsByProject(c *gin.Context)
	GetRecordById(c *gin.Context)
	CreateRecord(c *gin.Context)
	UpdateRecord(c *gin.Context)
	DeleteRecord(c *gin.Context)
}

/**
* expected format of json post/put requests
 **/
type PortfolioRecordInput struct {
	UserID       int
	Summary      string
	Technologies []string
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
* Create Record
 **/
func (rc portfolioRecordController) CreateRecord(c *gin.Context) {
	// Get POST data
	var input PortfolioRecordInput
	err := json.NewDecoder(c.Request.Body).Decode(&input)
	// print error if any
	if err != nil {
		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}

	// init New Record Model
	record := models.PortfolioRecord{
		Summary: input.Summary,
		UserID:  uint(input.UserID),
	}

	// get ProjectID from url
	projectID := c.Param("id")
	// get project
	var project models.Project
	err = database.DB.Where("id = ?", projectID).First(&project).Error
	if err != nil {
		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	} else {
		record.ProjectID = project.ID
	}

	// Technologies
	if input.Technologies != nil {
		var technologies []models.Technology
		for _, v := range input.Technologies {

			// TODO: Add check for lowercase/uppercase/caps
			var tech models.Technology
			database.DB.Where(models.Technology{Name: v}).FirstOrInit(&tech)
			technologies = append(technologies, tech)
		}
		record.Technologies = &technologies
	}

	// create record with repo
	rec, err := rc.recordRepository.CreateRecord(record)
	if err != nil {
		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(c.Writer, http.StatusOK, rec)
}

/**
* Update Record
 **/
func (rc portfolioRecordController) UpdateRecord(c *gin.Context) {
	// Get POST data
	var input PortfolioRecordInput
	err := json.NewDecoder(c.Request.Body).Decode(&input)
	// print error if any
	if err != nil {
		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}

	// get existing record
	id := c.Param("id")
	// get project
	record, err := rc.recordRepository.FindRecordById(id)
	if err != nil {
		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}

	// init New Record Model
	newRecordModel := models.PortfolioRecord{
		Summary: input.Summary,
		UserID:  uint(input.UserID),
	}

	// Technologies
	if input.Technologies != nil {
		var technologies []models.Technology
		for _, v := range input.Technologies {

			// TODO: Add check for lowercase/uppercase/caps
			var tech models.Technology
			database.DB.Where(models.Technology{Name: v}).FirstOrInit(&tech)
			technologies = append(technologies, tech)
		}
		newRecordModel.Technologies = &technologies
	}

	// Update Record
	rec, err := rc.recordRepository.UpdateRecord(record, newRecordModel)
	if err != nil {
		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(c.Writer, http.StatusOK, rec)
}

/**
* Delete Record
 **/
func (rc portfolioRecordController) DeleteRecord(c *gin.Context) {
	id := c.Param("id")
	err := rc.recordRepository.DeleteRecordById(id)
	if err != nil {
		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(c.Writer, http.StatusOK, "Portfolio Record Deleted Successfully")
}
