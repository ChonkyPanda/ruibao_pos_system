package service

import (
	"golang.org/x/crypto/bcrypt"

	"ruibao_pos_system/repository"
	"ruibao_pos_system/models"
)

func GetAllUsers() ([]models.User, error) {
 return repository.GetAllUsers()
}

func GetUserByUsername(username string) (*models.User, error) {
	return repository.GetUserByUsername(username)
}

func GetUserByID(id uint) (*models.User, error) {
	return repository.GetUserByID(id)
}

func CreateUser(name, password string, admin bool) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return repository.CreateUser(name, string(hashedPassword), admin)
}

func UpdateUser(u *models.User) error {
	return repository.UpdateUser(u)
}

func DeleteUser(id uint) error {
	return repository.DeleteUser(id)
}

