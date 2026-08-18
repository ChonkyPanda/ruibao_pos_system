package repository

import (
	"gorm.io/gorm"

	"ruibao_pos_system/database"
	"ruibao_pos_system/models"
)

func Query[T any](query *gorm.DB) (*T, error) {
	var res T
	err := query.First(&res).Error
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := database.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserByUsername(username string) (*models.User, error) {
	return Query[models.User](database.DB.Where("username = ?", username))
}

func GetUserByID(id uint) (*models.User, error) {
	return Query[models.User](database.DB.First(&models.User{}, id))
}

func CreateUser(username, password string, admin bool) (*models.User, error) {
	u := &models.User{
		Username: username,
		Password: password,
		Admin:  admin,
	}
	err := database.DB.Create(u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func UpdateUser(u *models.User) error {
	return database.DB.Model(&models.User{}).Where("id = ?", u.ID).Updates(u).Error
}

func DeleteUser(id uint) error {
	return database.DB.Delete(&models.User{}, id).Error
}