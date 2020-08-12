package main

import (
  "github.com/labstack/echo/middleware"
  "github.com/labstack/echo"
  "./handler"
)

func main() {
  e := echo.New()
  // ログなど
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  e.Use(middleware.CORS()) // CORS対策のはず

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
  
  e.GET("/answer/:id", handler.GetAnswer)  // :idはAnswerのIDではなくQuestionのID

  e.GET("/comment/:id", handler.GetComment)  // :idはAnswerのIDではなくQuestionのID

  e.Start(":1323")
}