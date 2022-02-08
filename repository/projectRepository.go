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
	FindProjectById(id string) (models.Project, error)
	CreateProject(project models.Project) (models.Project, error)
	UpdateProjectById(id string, updatedValues models.Project) error
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
	err = p.DB.Find(&projects).Error
	return projects, err
}

/**
* Get single Project by id
 **/
func (p projectRepository) FindProjectById(id string) (project models.Project, err error) {
	err = p.DB.Where("id = ?", id).First(&project).Error
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
func (p projectRepository) UpdateProjectById(id string, updatedValues models.Project) error {
	err := p.DB.Model(&models.Project{}).Where("id = ?", id).Updates(updatedValues).Error
	return err
}

/**
* Delete Project by id
 **/
func (p projectRepository) DeleteProjectById(id string) error {
	err := p.DB.Delete(&models.Project{}, id).Error
	return err
}
