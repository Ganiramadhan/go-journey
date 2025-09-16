package service

import (
	"go-journey/src/database"
	"go-journey/src/model"
	"log"
)

// GetAllUsers fetches all users from DB
func GetAllUsers() ([]model.User, error) {
	var users []model.User
	result := database.DB.Find(&users)
	return users, result.Error
}

// GetUserByID fetches a single user by UUID
func GetUserByID(id string) (model.User, error) {
	var user model.User
	result := database.DB.Where("id = ?", id).First(&user)
	return user, result.Error
}

// CreateUser inserts a new user into DB
func CreateUser(user *model.User) error {
	result := database.DB.Create(user)
	return result.Error
}

// UpdateUser updates existing user data
func UpdateUser(user *model.User) error {
	result := database.DB.Save(user)
	return result.Error
}

func DeleteUser(id string) error {
	result := database.DB.Delete(&model.User{}, "id = ?", id)
	if result.Error != nil {
		log.Println("[DeleteUser] DB Error:", result.Error)
	} else {
		log.Println("[DeleteUser] Rows Affected:", result.RowsAffected)
	}
	return result.Error
}
