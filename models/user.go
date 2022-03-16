package models

import (
	"echo_todolist/database"
	"echo_todolist/database_errors"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `json: "userName"`
	// Todolist []Todo
}

// Input of CreateUsers/ UpdateAUser
type UserInfo struct {
	UserName string `json: "userName"`
}

//func Show all users
func FindAllUsers() ([]User, error) {
	var users []User
	result := database.DB.Find(&users)
	if result.Error != nil {
		err := fmt.Errorf(database_errors.Err500)
		return nil, err
	}
	return users, nil
}

// func Create Users
func CreateUsers(userNew User) (userCreated User, err error) {
	result := database.DB.Create(&userNew)
	userCreated = userNew
	if result.Error != nil {
		err := fmt.Errorf(database_errors.Err500)
		return userCreated, err
		// chua tim duoc gia tri rong cua user
	}
	return userCreated, nil
}

func FindAUser(userID string) (user User, err error) {
	result := database.DB.First(&user, userID)
	isNotFoundError := errors.Is(result.Error, gorm.ErrRecordNotFound)
	if result.Error != nil {
		if isNotFoundError {
			err := fmt.Errorf(database_errors.Err404)
			return user, err
		} else {
			err := fmt.Errorf(database_errors.Err500)
			return user, err
		}
	}
	return user, nil
}

func UpdateAUser(userID string, userUpdate User) (User, error) {
	// find user need update
	user, err := FindAUser(userID)
	if err != nil {
		return user, err
	}
	// Update User
	result := database.DB.Model(&user).Updates(&userUpdate)
	if result.Error != nil {
		err := fmt.Errorf(database_errors.Err500)
		return user, err
	}
	return user, nil
}

func DeleteAUser(userID string) error {
	var user User
	result := database.DB.Delete(&user, userID)
	if result.Error != nil {
		err := fmt.Errorf(database_errors.Err500)
		return err
	}
	return nil
}
