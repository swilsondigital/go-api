package repository

import (
	"goapi/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type portfolioRecordRepository struct {
	DB *gorm.DB
}

type PortfolioRecordRepository interface {
	FindRecordsByProjectId(id string) (models.PortfolioRecords, error)
	FindRecordsByUserId(id string) (models.PortfolioRecords, error)
	FindRecordById(id string) (models.PortfolioRecord, error)
	CreateRecord(record models.PortfolioRecord) (models.PortfolioRecord, error)
	UpdateRecord(record models.PortfolioRecord, updatedValues models.PortfolioRecord) (models.PortfolioRecord, error)
	DeleteRecordById(id string) error
}

/**
* Get new PortfolioRecord repository instance
 **/
func NewPortfolioRecordRepository(db *gorm.DB) PortfolioRecordRepository {
	return portfolioRecordRepository{DB: db}
}

/**
* Get All Records for a project
 **/
func (p portfolioRecordRepository) FindRecordsByProjectId(id string) (records models.PortfolioRecords, err error) {
	err = p.DB.Where("project_id = ?", id).Preload("Project.Client").Preload(clause.Associations).Find(&records).Error
	return records, err
}

/**
* Get All Records for a user
 **/
func (p portfolioRecordRepository) FindRecordsByUserId(id string) (records models.PortfolioRecords, err error) {
	err = p.DB.Where("user_id = ?", id).Preload("Project.Client").Preload(clause.Associations).Find(&records).Error
	return records, err
}

/**
* Get Single Record
 **/
func (p portfolioRecordRepository) FindRecordById(id string) (record models.PortfolioRecord, err error) {
	err = p.DB.Where("id = ?", id).Preload("Project.Client").Preload(clause.Associations).First(&record).Error
	return record, err
}

/**
* Create New Record
 **/
func (p portfolioRecordRepository) CreateRecord(record models.PortfolioRecord) (models.PortfolioRecord, error) {
	err := p.DB.Create(&record).Error
	return record, err
}

/**
* Update Record
 **/
func (p portfolioRecordRepository) UpdateRecord(record models.PortfolioRecord, updatedValues models.PortfolioRecord) (models.PortfolioRecord, error) {
	// update record model
	err := p.DB.Model(&record).Updates(&updatedValues).Error

	// upsert technologies
	if updatedValues.Technologies != nil {
		p.DB.Model(&record).Association("Technologies").Replace(&updatedValues.Technologies)
	} else {
		p.DB.Model(&record).Association("Technologies").Clear()
	}

	// save to make sure everything persists
	p.DB.Save(&record)
	return record, err
}

/**
* Delete Record by ID
 **/
func (p portfolioRecordRepository) DeleteRecordById(id string) error {
	err := p.DB.Delete(&models.PortfolioRecord{}, id).Error
	return err
}
