package repository

import (
	"goapi/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepository struct {
	DB *gorm.DB
}

type UserRepository interface {
	FindAllUsers() (models.Users, error)
	FindUserById(id string) (models.User, error)
	FindUserByEmail(email string) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User, updatedValues models.User) (models.User, error)
	DeleteUserById(id string) error
}

/**
* Generate New User Repository
 **/
func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{DB: db}
}

/**
* Get all users
 **/
func (u userRepository) FindAllUsers() (users models.Users, err error) {
	err = u.DB.Preload(clause.Associations).Find(&users).Error
	return users, err
}

/**
* Get single user by id
 **/
func (u userRepository) FindUserById(id string) (user models.User, err error) {
	err = u.DB.Preload(clause.Associations).Preload("PortfolioRecords.Technologies").Where("id = ?", id).First(&user).Error
	return user, err
}

/**
* Get single user by id
 **/
func (u userRepository) FindUserByEmail(email string) (user models.User, err error) {
	err = u.DB.Preload(clause.Associations).Preload("PortfolioRecords.Technologies").Where("email = ?", email).First(&user).Error
	return user, err
}

/**
* Create a new user
 **/
func (u userRepository) CreateUser(user models.User) (models.User, error) {
	err := u.DB.Create(&user).Error
	return user, err
}

/**
* Update user
 **/
func (u userRepository) UpdateUser(user models.User, updatedValues models.User) (models.User, error) {
	// update base user model
	err := u.DB.Model(&user).Updates(updatedValues).Error

	// upsert Profile
	if updatedValues.Profile.UserID != 0 {
		// update Associated Profile
		u.DB.Model(&user).Association("Profile").Append(&updatedValues.Profile)
	} else {
		u.DB.Model(&user).Association("Profile").Clear()
	}

	// save to make sure everything persists
	u.DB.Save(&user)
	return user, err
}

/**
* Delete user by ID
 **/
func (u userRepository) DeleteUserById(id string) error {
	err := u.DB.Delete(&models.User{}, id).Error
	return err
}
