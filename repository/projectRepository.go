package repository

import (
	"goapi/models"

	"gorm.io/gorm"
)

type projectRepository struct {
	DB *gorm.DB
}

type ProjectRepository interface {
	FindAllProjects() (models.Projects, error)
	FindAllProjectsByClientId(id string) (models.Projects, error)
	FindProjectById(id string) (models.Project, error)
	CreateProject(project models.Project) (models.Project, error)
	UpdateProject(project models.Project, updatedValues models.Project) (models.Project, error)
	DeleteProjectById(id string) error
}

/**
* Get new Project repository instance
 **/
func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return projectRepository{DB: db}
}

/**
* Get all Projects
 **/
func (p projectRepository) FindAllProjects() (projects models.Projects, err error) {
	err = p.DB.Preload("Client.Logo").Find(&projects).Error
	return projects, err
}

/**
* Get All Projects By Client ID
 **/
func (p projectRepository) FindAllProjectsByClientId(id string) (projects models.Projects, err error) {
	err = p.DB.Preload("Client.Logo").Preload("Technologies").Where("client_id = ?", id).Find(&projects).Error
	return projects, err
}

/**
* Get single Project by id
 **/
func (p projectRepository) FindProjectById(id string) (project models.Project, err error) {
	err = p.DB.Preload("Client.Logo").Preload("Technologies").Where("id = ?", id).First(&project).Error
	return project, err
}

/**
* Create new Project
 **/
func (p projectRepository) CreateProject(project models.Project) (models.Project, error) {
	err := p.DB.Create(&project).Error
	return project, err
}

/**
* Update Project by id
 **/
func (p projectRepository) UpdateProject(project models.Project, updatedValues models.Project) (models.Project, error) {
	err := p.DB.Model(&project).Updates(updatedValues).Error

	// upsert technologies
	if updatedValues.Technologies != nil {
		p.DB.Model(&project).Association("Technologies").Replace(&updatedValues.Technologies)
	} else {
		p.DB.Model(&project).Association("Technologies").Clear()
	}

	// save to make sure everything persists
	p.DB.Save(&project)
	return project, err
}

/**
* Delete Project by id
 **/
func (p projectRepository) DeleteProjectById(id string) error {
	err := p.DB.Delete(&models.Project{}, id).Error
	return err
}
