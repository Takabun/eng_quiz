package handler

import (
	"net/http"
	"github.com/labstack/echo"
	. "../controller"
	. "../model"
)

func CreateQuestionImage(c echo.Context) error {
  db := OpenSQLiteConnection()
  defer db.Close()
  db.AutoMigrate(&QuestionImage{})

  questionimage := new(QuestionImage)
  if err := c.Bind(questionimage); err != nil {
    return err
  }
	db.Create(&questionimage)

  return c.JSON(http.StatusOK, questionimage)
}

func DeleteQuestionImage(c echo.Context) error {
  db := OpenSQLiteConnection()
  defer db.Close()

  if id := c.Param("id"); id != "" {
    var questionimage QuestionImage
    db.First(&questionimage, id)
    db.Delete(questionimage)
    return c.JSON(http.StatusOK, questionimage)
  } else {
    return c.JSON(http.StatusNotFound, nil)
  }
}