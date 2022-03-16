package controllers

import (
	"echo_todolist/models"
)

func GetUsers() ([]models.User, error) {
	users, err := models.FindAllUsers()
	return users, err
}

func CreateUsers(createUserInfo models.UserInfo) (user models.User, err error) {
	user.UserName = createUserInfo.UserName
	user, err = models.CreateUsers(user)
	return user, err
}

func GetAUser(userID string) (user models.User, err error) {
	user, err = models.FindAUser(userID)
	return user, err
}

func UpdateAUser(userID string, updateUserInfo models.UserInfo) (user models.User, err error) {
	user.UserName = updateUserInfo.UserName
	user, err = models.UpdateAUser(userID, user)
	return user, err
}

func DeleteAUser(userID string) error {
	err := models.DeleteAUser(userID)
	return err
}
