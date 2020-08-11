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
  e.GET("/questions", handler.GetAllQuestions)
  e.GET("/question/:id", handler.GetQuestion)
  e.POST("/question", handler.CreateQuestion)
  e.PUT("/question/:id", handler.UpdateQuestion)
  e.DELETE("/question/:id", handler.DeleteQuestion)

  e.GET("/users", handler.GetAllUsers)

  
  e.Start(":1323")
}