package views

import (
	"echo_todolist/controllers"
	"echo_todolist/database_errors"
	"echo_todolist/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	users, err := controllers.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}

func CreateUsers(c echo.Context) error {
	var createUserInfo models.UserInfo
	if err := c.Bind(&createUserInfo); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	user, err := controllers.CreateUsers(createUserInfo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, user)
}

func GetAUser(c echo.Context) error {
	userID := c.Param("userID")
	user, err := controllers.GetAUser(userID)
	if err != nil {
		if err.Error() == database_errors.Err500 {
			return c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			return c.JSON(http.StatusNotFound, err.Error())
		}
	}
	return c.JSON(http.StatusOK, user)
}

func UpdateAUser(c echo.Context) error {
	userID := c.Param("userID")
	var updateUserInfo models.UserInfo
	if err := c.Bind(&updateUserInfo); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	user, err := controllers.UpdateAUser(userID, updateUserInfo)
	if err != nil {
		if err.Error() == database_errors.Err500 {
			return c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			return c.JSON(http.StatusNotFound, err.Error())
		}
	}
	return c.JSON(http.StatusOK, user)
}

func DeleteAUser(c echo.Context) error {
	userID := c.Param("userID")
	err := controllers.DeleteAUser(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Delete Successful")
}
