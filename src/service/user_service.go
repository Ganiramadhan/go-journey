package service

import (
	"go-journey/src/database"
	"go-journey/src/model"
)

// Get all users
func GetAllUsers() ([]model.User, error) {
	var users []model.User
	result := database.DB.Find(&users)
	return users, result.Error
}

// Get user by ID
func GetUserByID(id uint) (model.User, error) {
	var user model.User
	result := database.DB.First(&user, id)
	return user, result.Error
}

// Create new user
func CreateUser(user *model.User) error {
	result := database.DB.Create(user)
	return result.Error
}

// Update user
func UpdateUser(user *model.User) error {
	result := database.DB.Save(user)
	return result.Error
}

// Delete user
func DeleteUser(id uint) error {
	result := database.DB.Delete(&model.User{}, id)
	return result.Error
}
