package repository

import (
	"goapi/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type clientRepository struct {
	DB *gorm.DB
}

type ClientRepository interface {
	FindAllClients() (models.Clients, error)
	FindClientById(id string) (models.Client, error)
	CreateClient(client models.Client) (models.Client, error)
	UpdateClient(client models.Client, updatedValues models.Client) (models.Client, error)
	DeleteClientById(id string) error
}

/**
* Get new client repository instance
 **/
func NewClientRepository(db *gorm.DB) ClientRepository {
	return clientRepository{DB: db}
}

/**
* Get all clients
 **/
func (c clientRepository) FindAllClients() (clients models.Clients, err error) {
	err = c.DB.Preload(clause.Associations).Preload("Projects.Technologies").Preload("Projects.PortfolioRecords.User").Find(&clients).Error
	return clients, err
}

/**
* Get single client by id
 **/
func (c clientRepository) FindClientById(id string) (client models.Client, err error) {
	err = c.DB.Preload(clause.Associations).Preload("Projects.Technologies").Preload("Projects.PortfolioRecords.User").Where("id = ?", id).First(&client).Error
	return client, err
}

/**
* Create new client
 **/
func (c clientRepository) CreateClient(client models.Client) (models.Client, error) {
	err := c.DB.Create(&client).Error
	return client, err
}

/**
* Update client by id
 **/
func (c clientRepository) UpdateClient(client models.Client, updatedValues models.Client) (models.Client, error) {
	// update base client model
	err := c.DB.Model(&client).Updates(updatedValues).Error

	// upsert Address
	c.DB.Model(&client).Association("Address").Append(&updatedValues.Address)

	// upsert Logo/Image
	if updatedValues.Logo.Blob == "" {
		// remove if set to empty string -> check to see if this deletes the associated db entry
		c.DB.Model(&client).Association("Logo").Clear()
	} else {
		// replace existing logo association -> maybe update this to replace string in association
		c.DB.Model(&client).Association("Logo").Replace(&updatedValues.Logo)
	}

	// upsert Contact/User
	if updatedValues.Contact.ID != 0 {
		// update client contact
		c.DB.Model(&client).Association("Contact").Append(&updatedValues.Contact)
	} else {
		// remove client contact
		c.DB.Model(&client).Association("Contact").Clear()
	}

	// save to make sure everything persists
	c.DB.Save(&client)
	return client, err
}

/**
* Delete client by id
 **/
func (c clientRepository) DeleteClientById(id string) error {
	err := c.DB.Delete(&models.Client{}, id).Error
	return err
}
