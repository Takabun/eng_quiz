package handler

import (
	"net/http"
	"github.com/labstack/echo"
	. "../controller"
	. "../model"
)

func CreateTag(c echo.Context) error {
  db := OpenSQLiteConnection()
  defer db.Close()
  db.AutoMigrate(&Tag{})

  tag := new(Tag)
  if err := c.Bind(tag); err != nil {
    return err
  }
  db.Create(&tag)

  return c.JSON(http.StatusOK, tag)
}