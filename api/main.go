package main

import (
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  "./handler"
)

func main() {
  e := echo.New()
  // ログなど
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())
  // ルーティング
  initRouting(e) 
}

func initRouting(e *echo.Echo) {
	// e.GET("/high_school/testcrate", handler.CreateMember)  //ユーザー作成
  // e.GET("/high_school/:name", handler.GetMember)
  // e.GET("/high_school/all", handler.GetAllMembers)
  e.GET("/posts", handler.GetAllQuestions)
  e.GET("/posts/:id", handler.GetQuestion)
  e.POST("/posts", handler.CreateQuestion)
  e.PUT("/posts/:id", handler.UpdateQuestion)
  e.DELETE("/posts/:id", handler.DeleteQuestion)
  e.Start(":1323")
}