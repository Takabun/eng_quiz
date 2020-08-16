package handler

import (
	"net/http"
	"github.com/labstack/echo"
	. "../controller"
	. "../model"
)

// 個別のAnswer(:idはAnswerのIDではなくQuestionのID)
func GetAnswer(c echo.Context) error {
	db := OpenSQLiteConnection()
	defer db.Close()
	db.AutoMigrate(&Answer{})

	if id := c.Param("id"); id != "" {
		var answer Answer
		db.Where("question_id = ?", id).First(&answer)
		db.Model(&answer).Related(&answer.AnswerImage) //上に繋げたら動かなかったけど、こうしたら動いた
		return c.JSON(http.StatusOK, answer)
	} else {
		return c.JSON(http.StatusNotFound, nil)
	}
}


func CreateAnswer(c echo.Context) error {
  db := OpenSQLiteConnection()
  defer db.Close()
  db.AutoMigrate(&Answer{})

  answer := new(Answer)
  if err := c.Bind(answer); err != nil {
    return err
  }
	db.Create(&answer)
	
	// ここから画像をAssociation
	// db.Model(&answer).Association("AnswerImage").Append([]AnswerImage {{Url: "0816", Name: "0816"}, {Url: "0816", Name: "0816"}})

	// keke := c.FormValue("keke")
	// images := []AnswerImage{keke} 
	// // images := []AnswerImage{AnswerImage keke}
	// db.Model(&answer).Association("AnswerImage").Append(images)

  return c.JSON(http.StatusOK, answer)
}


func UpdateAnswer(c echo.Context) error {
  db := OpenSQLiteConnection()
  defer db.Close()

  newAnswer := new(Answer)
  if err := c.Bind(newAnswer); err != nil {
    return err
  }

  if id := c.Param("id"); id != "" {
    var answer Answer
    db.First(&answer, id).Update(newAnswer)
    return c.JSON(http.StatusOK, answer)
  } else {
    return c.JSON(http.StatusNotFound, nil)
  }
}

func DeleteAnswer(c echo.Context) error {
  db := OpenSQLiteConnection()
  defer db.Close()

  if id := c.Param("id"); id != "" {
    var answer Answer
    db.First(&answer, id)
    db.Delete(answer)
    return c.JSON(http.StatusOK, answer)
  } else {
    return c.JSON(http.StatusNotFound, nil)
  }
}