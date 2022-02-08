package repository

import (
	"goapi/models"

	"gorm.io/gorm"
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
	err = u.DB.Find(&users).Error
	return users, err
}

/**
* Get single user by id
 **/
func (u userRepository) FindUserById(id string) (user models.User, err error) {
	err = u.DB.Where("id = ?", id).First(&user).Error
	return user, err
}

/**
* Get single user by id
 **/
func (u userRepository) FindUserByEmail(email string) (user models.User, err error) {
	err = u.DB.Where("email = ?", email).First(&user).Error
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
	if updatedValues.Profile.UserID == user.ID {
		// update Associated Profile
		u.DB.Model(&user).Association("Profile").Replace(&updatedValues.Profile)
	} else {
		u.DB.Model(&user).Association("Profile").Append(&updatedValues.Profile)
	}

	return user, err
}

/**
* Delete user by ID
 **/
func (u userRepository) DeleteUserById(id string) error {
	err := u.DB.Delete(&models.User{}, id).Error
	return err
}
