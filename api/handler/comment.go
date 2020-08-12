package handler

import (
	"net/http"
	"github.com/labstack/echo"
	. "../controller"
	. "../model"
)


// 個別のComment(:idはCommentのIDではなくQuestionのID)
func GetComment(c echo.Context) error {
	db := OpenSQLiteConnection()
	defer db.Close()
	db.AutoMigrate(&Comment{})

	if id := c.Param("id"); id != "" {
		var comments []Comment  // Commentは複数あるので[]
		db.Find(&comments).Where("question_id = ?", id)

		for i := range comments {
			db.Model(comments[i]).Related(&comments[i].User) // 第二引数("Users")なくてもイケた
		}

		return c.JSON(http.StatusOK, comments)
	} else {
		return c.JSON(http.StatusNotFound, nil)
	}
}
