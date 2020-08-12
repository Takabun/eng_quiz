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
