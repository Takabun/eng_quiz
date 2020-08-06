package handler

import (
    "net/http"
    "github.com/labstack/echo"
    . "../controller"
)

type Question struct {
    Id int `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
    Text string `json:"text"`
}

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
        db.First(&question, id)
        return c.JSON(http.StatusOK, question)
    } else {
        return c.JSON(http.StatusNotFound, nil)
    }
}

func CreateQuestion(c echo.Context) error {
    db := OpenSQLiteConnection()
    defer db.Close()
    db.AutoMigrate(&Question{})

    // question := new(Question)
    question := new(Question)
    if err := c.Bind(question); err != nil {
        return err
    }
    db.Create(&question)

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