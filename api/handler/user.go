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


func GetUser(c echo.Context) error {
  db := OpenSQLiteConnection()
  defer db.Close()
  db.AutoMigrate(&User{})

  if id := c.Param("id"); id != "" {
    var user User
    db.First(&user, id)
    return c.JSON(http.StatusOK, user)
  } else {
    return c.JSON(http.StatusNotFound, nil)
  }
}


func CreateUser(c echo.Context) error {
  db := OpenSQLiteConnection()
  defer db.Close()
  db.AutoMigrate(&User{})

  user := new(User)
  if err := c.Bind(user); err != nil {
    return err
  }
  db.Create(&user)

  return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
  db := OpenSQLiteConnection()
  defer db.Close()

  newUser := new(User)
  if err := c.Bind(newUser); err != nil {
    return err
  }

  if id := c.Param("id"); id != "" {
    var user User
    db.First(&user, id).Update(newUser)
    return c.JSON(http.StatusOK, user)
  } else {
    return c.JSON(http.StatusNotFound, nil)
  }
}

func DeleteUser(c echo.Context) error {
  db := OpenSQLiteConnection()
  defer db.Close()

  if id := c.Param("id"); id != "" {
    var user User
    db.First(&user, id)
    db.Delete(user)
    return c.JSON(http.StatusOK, user)
  } else {
    return c.JSON(http.StatusNotFound, nil)
  }
}