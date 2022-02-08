package repository

import (
	"goapi/models"

	"gorm.io/gorm"
)

type clientRepository struct {
	DB *gorm.DB
}

type ClientRepository interface {
	FindAllClients() (models.Clients, error)
	FindClientById(id string) (models.Client, error)
	CreateClient(client models.Client) (models.Client, error)
	UpdateClientById(id string, updatedValues models.Client) error
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
	err = c.DB.Find(&clients).Error
	return clients, err
}

/**
* Get single client by id
 **/
func (c clientRepository) FindClientById(id string) (client models.Client, err error) {
	err = c.DB.Where("id = ?", id).First(&client).Error
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
func (c clientRepository) UpdateClientById(id string, updatedValues models.Client) error {
	err := c.DB.Model(&models.Client{}).Where("id = ?", id).Updates(updatedValues).Error
	return err
}

/**
* Delete client by id
 **/
func (c clientRepository) DeleteClientById(id string) error {
	err := c.DB.Delete(&models.Client{}, id).Error
	return err
}
