package main

import (
	"errors"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

type Todo struct {
	gorm.Model
	TodoContent string `json: "todoContent"`
	isDone      bool   `json: "isDone"`
	UserID      uint32 `json: "userID"`
}
type User struct {
	gorm.Model
	UserName  string `json: "userName"`
	Todolists []Todo
}

func main() {
	db, err := gorm.Open(sqlite.Open("todolist.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Todo{}, &User{})

	e := echo.New()
	// Khoi tao 1 echo

	// CRUD User
	// Show all users
	e.GET("/users", func(c echo.Context) error {
		var users []User
		result := db.Find(&users)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		return c.JSON(http.StatusOK, users)
	})

	// Create a user
	e.POST("/users", func(c echo.Context) error {
		var user User
		if err := c.Bind(&user); err != nil { // ghi du lieu nguoi dung nhap vao vao bien user
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		result := db.Create(&user)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, result.Error)
		}
		return c.JSON(http.StatusCreated, user)
	})

	// show user information by userID
	e.GET("/users/:userID", func(c echo.Context) error {
		var user User
		userID := c.Param("userID") // get UserID on URL
		result := db.First(&user, userID)
		isNotFoundError := errors.Is(result.Error, gorm.ErrRecordNotFound)
		if result.Error != nil {
			if isNotFoundError {
				return c.JSON(http.StatusNotFound, "UserID does not exist")
			} else {
				return c.JSON(http.StatusInternalServerError, "Internal Server Error")
			}
		}
		return c.JSON(http.StatusOK, user)
	})

	// update user's information
	e.PATCH("users/:userID", func(c echo.Context) error {
		var user User
		userID := c.Param("userID")
		resultCheckExist := db.First(&user, userID)
		isNotFoundError := errors.Is(resultCheckExist.Error, gorm.ErrRecordNotFound)
		if resultCheckExist.Error != nil {
			if isNotFoundError {
				return c.JSON(http.StatusNotFound, "UserID does not exist")
			} else {
				return c.JSON(http.StatusInternalServerError, "Internal Server Error")
			}
		}
		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		resultUpdate := db.Model(&user).Updates(&user)
		if resultUpdate.Error != nil {
			return c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		return c.JSON(http.StatusOK, user)
	})
	// delete user
	e.DELETE("/users/:userID", func(c echo.Context) error {
		var user User
		userID := c.Param("userID")
		result := db.Delete(&user, userID)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		return c.JSON(http.StatusOK, "Delete Successful")
	})

	//CRUD todo
	// e.GET("users/:userID/todos", func(c echo.Context) error {
	// 	var todos []Todo
	// 	userID := c.Param("userID")

	// })

	e.Logger.Fatal(e.Start(":1323"))
}
