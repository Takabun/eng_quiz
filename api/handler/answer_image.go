package handler

import (
	"net/http"
	"github.com/labstack/echo"
	. "../controller"
	. "../model"
)

func CreateAnswerImage(c echo.Context) error {
  db := OpenSQLiteConnection()
  defer db.Close()
  db.AutoMigrate(&AnswerImage{})

  answerimage := new(AnswerImage)
  if err := c.Bind(answerimage); err != nil {
    return err
  }
	db.Create(&answerimage)

  return c.JSON(http.StatusOK, answerimage)
}


func DeleteAnswerImage(c echo.Context) error {
  db := OpenSQLiteConnection()
  defer db.Close()

  if id := c.Param("id"); id != "" {
    var answerimage AnswerImage
    db.First(&answerimage, id)
    db.Delete(answerimage)
    return c.JSON(http.StatusOK, answerimage)
  } else {
    return c.JSON(http.StatusNotFound, nil)
  }
}