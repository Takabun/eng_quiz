package handler

import (
	"net/http"
	"github.com/labstack/echo"
	. "../controller"
	. "../model"
)

func GetAllUsers(c echo.Context) error {
	db := OpenSQLiteConnection()
	defer db.Close()
	db.AutoMigrate(&User{})

	var users []User
	db.Find(&users)
	return c.JSON(http.StatusOK, users)
}
