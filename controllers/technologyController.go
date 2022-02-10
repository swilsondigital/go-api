package controllers

import (
	"goapi/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type technologyController struct {
	// place for auth later
	technologyRepository repository.TechnologyRepository
}
type TechnologyController interface {
	GetAllTechnologies(c *gin.Context)
	GetTechnologyById(c *gin.Context)
	GetTechnologyByName(c *gin.Context)
}

/**
* expected format of json post/put requests
 **/
type TechnologyInput struct {
	Name string
}

/**
* Setup New Technology Controller
 **/
func NewTechnologyController(tr repository.TechnologyRepository) TechnologyController {
	return technologyController{technologyRepository: tr}
}

/**
* Get All Technologies
 **/
func (tc technologyController) GetAllTechnologies(c *gin.Context) {
	technologies, err := tc.technologyRepository.FindAllTechnologies()
	if err != nil {
		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(c.Writer, http.StatusOK, technologies)
}

/**
* Get Technology by id
 **/
func (tc technologyController) GetTechnologyById(c *gin.Context) {
	id := c.Param("id")
	tech, err := tc.technologyRepository.FindTechnologyById(id)
	if err != nil {
		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(c.Writer, http.StatusOK, tech)
}

/**
* Get Technology by name
 **/
func (tc technologyController) GetTechnologyByName(c *gin.Context) {
	name := c.Query("name")
	// TODO: Add check for lowercase/uppercase/caps
	tech, err := tc.technologyRepository.FindTechnologyByName(name)
	if err != nil {
		RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(c.Writer, http.StatusOK, tech)
}
