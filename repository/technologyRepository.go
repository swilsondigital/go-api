package repository

import (
	"goapi/models"

	"gorm.io/gorm"
)

type technologyRepository struct {
	DB *gorm.DB
}

type TechnologyRepository interface {
	FindAllTechnologies() (models.Technologies, error)
	FindTechnologyById(id string) (models.Technology, error)
	FindTechnologyByName(name string) (models.Technology, error)
	CreateTechnology(technology models.Technology) (models.Technology, error)
	UpdateTechnologyById(id string, updatedValues models.Technology) error
	DeleteTechnologyById(id string) error
}

/**
* Generate New Technology Repository
 **/
func NewTechnologyRepository(db *gorm.DB) TechnologyRepository {
	return technologyRepository{DB: db}
}

/**
* Get all technologies
 **/
func (t technologyRepository) FindAllTechnologies() (technologies models.Technologies, err error) {
	err = t.DB.Find(&technologies).Error
	return technologies, err
}

/**
* Get single technology by id
 **/
func (t technologyRepository) FindTechnologyById(id string) (technology models.Technology, err error) {
	err = t.DB.Where("id = ?", id).First(&technology).Error
	return technology, err
}

/**
* Get single technology by name
 **/
func (t technologyRepository) FindTechnologyByName(name string) (technology models.Technology, err error) {
	err = t.DB.Where("name = ?", name).First(&technology).Error
	return technology, err
}

/**
* Create a new technology
 **/
func (t technologyRepository) CreateTechnology(technology models.Technology) (models.Technology, error) {
	err := t.DB.Create(&technology).Error
	return technology, err
}

/**
* Update technology by ID
 **/
func (t technologyRepository) UpdateTechnologyById(id string, updatedValues models.Technology) error {
	err := t.DB.Model(&models.Technology{}).Where("id = ?", id).Updates(updatedValues).Error
	return err
}

/**
* Delete technology by ID
 **/
func (t technologyRepository) DeleteTechnologyById(id string) error {
	err := t.DB.Delete(&models.Technology{}, id).Error
	return err
}
