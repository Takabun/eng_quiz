package handler

import (
	"net/http"
	"github.com/labstack/echo"
	. "../controller"
	. "../model"
)

// 各ユーザー事のList(:idはMylistのIDではなくUserのID)
func GetMylistForUser(c echo.Context) error {
	db := OpenSQLiteConnection()
	defer db.Close()
	db.AutoMigrate(&Mylist{})

	if id := c.Param("id"); id != "" {
		var mylist []Mylist
		db.Where("user_id = ?", id).Find(&mylist)
		// db.Model(&mylist).Related(&mylist.Question) //上に繋げたら動かなかったけど、こうしたら動いた
		for i := range mylist {
			db.Model(mylist[i]).Related(&mylist[i].Question) // 第二引数("User")なくてもイケた
		}
		return c.JSON(http.StatusOK, mylist)
	} else {
		return c.JSON(http.StatusNotFound, nil)
	}
}

func CreateMylist(c echo.Context) error {
  db := OpenSQLiteConnection()
  defer db.Close()
  db.AutoMigrate(&Mylist{})

  mylist := new(Mylist)
  if err := c.Bind(mylist); err != nil {
    return err
  }
  db.Create(&mylist)

  return c.JSON(http.StatusOK, mylist)
}

func DeleteMylist(c echo.Context) error {
  db := OpenSQLiteConnection()
  defer db.Close()

  if id := c.Param("id"); id != "" {
    var mylist Mylist
    db.First(&mylist, id)
    db.Delete(mylist)
    return c.JSON(http.StatusOK, mylist)
  } else {
    return c.JSON(http.StatusNotFound, nil)
  }
}