package handler

import (
  "net/http"
  "github.com/labstack/echo"
  . "../controller"
  . "../model"
  // "encoding/json"
  // "log"
  // "fmt"

)

func GetAllQuestions(c echo.Context) error {
  db := OpenSQLiteConnection()
  defer db.Close()
  db.AutoMigrate(&Question{})

  var questions []Question
  db.Find(&questions)
  return c.JSON(http.StatusOK, questions)
}

func GetQuestion(c echo.Context) error {
  db := OpenSQLiteConnection()
  defer db.Close()
  db.AutoMigrate(&Question{})

  if id := c.Param("id"); id != "" {
    var question Question
    db.First(&question, id).Related(&question.User).Related(&question.Tags, "Tags").Related(&question.QuestionImage) // .Related(&question.Answer).Related(&question.Comments)
    return c.JSON(http.StatusOK, question)
  } else {
    return c.JSON(http.StatusNotFound, nil)
  }
}

func CreateQuestion(c echo.Context, ) error {
  db := OpenSQLiteConnection()
  defer db.Close()
  db.AutoMigrate(&Question{})

  question := new(Question)
  if err := c.Bind(question); err != nil {
    return err
  }
  db.Create(&question)

  // 今、これを動かせるようにトライ中。1つだけであれば動かせていたが…
  for i := range uint[]{1,2} {
    var tag Tag
    tag.ID = i
    db.Model(&question).Association("Tags").Append(&tag)
  }

  // そもそもc.Paramではダメ(/image/:idみたいなものでは無いので)
  // []int{1,2}とすれば取り敢えず動く | addtagids := c.Param("addtagids")
  
  // for i := range addtagids {  // birds
  //   var tag Tag
  //   tag.ID = uint(i) //uintへ変換しないと
  //   db.Model(&question).Association("Tags").Append(&tag)
  // }

  return c.JSON(http.StatusOK, question)
}

func UpdateQuestion(c echo.Context) error {
  db := OpenSQLiteConnection()
  defer db.Close()

  newQuestion := new(Question)
  if err := c.Bind(newQuestion); err != nil {
    return err
  }

  if id := c.Param("id"); id != "" {
    var question Question
    db.First(&question, id).Update(newQuestion)
    return c.JSON(http.StatusOK, question)
  } else {
    return c.JSON(http.StatusNotFound, nil)
  }
}

func DeleteQuestion(c echo.Context) error {
  db := OpenSQLiteConnection()
  defer db.Close()

  if id := c.Param("id"); id != "" {
    var question Question
    db.First(&question, id)
    db.Delete(question)
    return c.JSON(http.StatusOK, question)
  } else {
    return c.JSON(http.StatusNotFound, nil)
  }
}