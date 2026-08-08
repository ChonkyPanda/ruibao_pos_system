package db

import "gorm.io/gorm"
import "golang.org/x/crypto/bcrypt"

type User struct {
	gorm.Model
	Name  string
	Password string
	Admin bool
}

func CreateUser(name, password string, admin bool) (*User, error) {
    // Hash the password first
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    // Store the hashed string instead of raw password
    user := User{
        Name:     name,
        Password: string(hashedPassword),
        Admin:    admin,
    }

    result := DB.Create(&user)
    if result.Error != nil {
        return nil, result.Error
    }

    return &user, nil
}

func GetFirstUser() (*User, error) {
	var user User
	result := DB.First(&user)

	if result.Error != nil {
        return nil, result.Error
    }
	
    return &user, nil
}

func GetUserByName(name string) (*User, error) {
	var user User
	result := DB.Where("name = ?", name).First(&user)

	if result.Error != nil {
        return nil, result.Error
    }

    return &user, nil
}