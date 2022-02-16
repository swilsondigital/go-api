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

type projectController struct {
	// place for auth later
	projectRepository repository.ProjectRepository
}
type ProjectController interface {
	GetAllProjects(c *gin.Context)
	GetAllClientProjects(c *gin.Context)
	GetProjectById(c *gin.Context)
	CreateProject(c *gin.Context)
	UpdateProject(c *gin.Context)
	DeleteProject(c *gin.Context)
}

/**
* expected format of json post/put requests
 **/
type ProjectInput struct {
	Name          string
	Technologies  []string
	Start_Date    string // accepts yyyy-mm-dd
	Delivery_Date string // accepts yyyy-mm-dd
	Private       bool
}

/**
* Setup New Project Controller
 **/
func NewProjectController(pr repository.ProjectRepository) ProjectController {
	return projectController{projectRepository: pr}
}

/**
* Get All Projects
 **/
func (pc projectController) GetAllProjects(c *gin.Context) {
	projects, err := pc.projectRepository.FindAllProjects()
	if err != nil {
		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(c.Writer, http.StatusOK, projects)
}

/**
* Get All Client Projects
 **/
func (pc projectController) GetAllClientProjects(c *gin.Context) {
	id := c.Param("id")
	projects, err := pc.projectRepository.FindAllProjectsByClientId(id)
	if err != nil {
		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(c.Writer, http.StatusOK, projects)
}

/**
* Get Project by ID
 **/
func (pc projectController) GetProjectById(c *gin.Context) {
	id := c.Param("id")
	project, err := pc.projectRepository.FindProjectById(id)
	if err != nil {
		RespondWithError(c.Writer, http.StatusNotFound, err.Error())
		return
	}
	RespondWithJson(c.Writer, http.StatusOK, project)
}

/**
* Create New Project
 **/
func (pc projectController) CreateProject(c *gin.Context) {
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
		Start_Date:    &startDate,
		Delivery_Date: &deliveryDate,
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
	p, err := pc.projectRepository.CreateProject(project)
	if err != nil {
		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(c.Writer, http.StatusOK, p)
}

/**
* Update Project
 **/
func (pc projectController) UpdateProject(c *gin.Context) {
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
	project, err := pc.projectRepository.FindProjectById(id)
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
		Start_Date:    &startDate,
		Delivery_Date: &deliveryDate,
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
	p, err := pc.projectRepository.UpdateProject(project, newProjectModel)
	if err != nil {
		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(c.Writer, http.StatusOK, p)
}

/**
* Delete Project
 **/
func (pc projectController) DeleteProject(c *gin.Context) {
	id := c.Param("id")
	err := pc.projectRepository.DeleteProjectById(id)
	if err != nil {
		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(c.Writer, http.StatusOK, "Project Deleted Successfully")
}
