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

		// for i := range comments {
		// 	db.Model(comments[i]).Related(&comments[i].User) // 第二引数("User")なくてもイケた
		// }

		return c.JSON(http.StatusOK, comments)
	} else {
		return c.JSON(http.StatusNotFound, nil)
	}
}


func GetAllComments(c echo.Context) error {
	db := OpenSQLiteConnection()
	defer db.Close()
	db.AutoMigrate(&Comment{})

	var comments []Comment
	db.Find(&comments)
	return c.JSON(http.StatusOK, comments)
}

func CreateComment(c echo.Context) error {
  db := OpenSQLiteConnection()
  defer db.Close()
  db.AutoMigrate(&Comment{})

  comment := new(Comment)
  if err := c.Bind(comment); err != nil {
    return err
  }
  db.Create(&comment)

  return c.JSON(http.StatusOK, comment)
}


func UpdateComment(c echo.Context) error {
  db := OpenSQLiteConnection()
  defer db.Close()

  newComment := new(Comment)
  if err := c.Bind(newComment); err != nil {
    return err
  }

  if id := c.Param("id"); id != "" {
    var comment Comment
    db.First(&comment, id).Update(newComment)
    return c.JSON(http.StatusOK, comment)
  } else {
    return c.JSON(http.StatusNotFound, nil)
  }
}

func DeleteComment(c echo.Context) error {
  db := OpenSQLiteConnection()
  defer db.Close()

  if id := c.Param("id"); id != "" {
    var comment Comment
    db.First(&comment, id)
    db.Delete(comment)
    return c.JSON(http.StatusOK, comment)
  } else {
    return c.JSON(http.StatusNotFound, nil)
  }
}