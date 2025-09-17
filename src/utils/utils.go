package utils

import (
	"go-journey/src/database"
	"go-journey/src/model"
)

func SaveRefreshToken(userID string, refreshToken string) error {
	return database.DB.Model(&model.User{}).
		Where("id = ?", userID).
		Update("refresh_token", refreshToken).Error
}

func IsRefreshTokenValid(userID string, refreshToken string) bool {
	var user model.User
	if err := database.DB.Select("refresh_token").
		Where("id = ?", userID).First(&user).Error; err != nil {
		return false
	}
	return user.RefreshToken == refreshToken
}

func RevokeRefreshToken(userID string) error {
	return database.DB.Model(&model.User{}).
		Where("id = ?", userID).
		Update("refresh_token", "").Error
}
