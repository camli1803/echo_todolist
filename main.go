package main

import (
	"echo_todolist/database"
	"echo_todolist/models"
	"echo_todolist/views"

	"github.com/labstack/echo/v4"
)

// type Todo struct {
// 	gorm.Model
// 	TodoContent string `json: "todoContent"`
// 	isDone      bool   `json: "isDone"`
// 	UserID      uint32 `json: "userID"`
// }
// type User struct {
// 	gorm.Model
// 	UserName string `json: "userName"`
// 	// Todolist []Todo
// }

func main() {
	// db, err := gorm.Open(sqlite.Open("todolist.db"), &gorm.Config{})
	db, err := database.ConnectDB()
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate( //&Todo{},
		&models.User{})

	e := echo.New()
	// Khoi tao 1 echo

	// CRUD User
	// Show all users
	e.GET("/users", views.GetUsers)

	// Create a user
	e.POST("/users", views.CreateUsers)

	// show user information by userID
	e.GET("/users/:userID", views.GetAUser)

	// update user's information
	e.PATCH("users/:userID", views.UpdateAUser)

	// delete user
	e.DELETE("/users/:userID", views.DeleteAUser)

	//CRUD todo
	// e.GET("users/:userID/todos", func(c echo.Context) error {
	// 	var todos []Todo
	// 	userID := c.Param("userID")

	// })

	e.Logger.Fatal(e.Start(":1323"))
}
