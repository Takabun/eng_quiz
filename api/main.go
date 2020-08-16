package main

import (
  "github.com/labstack/echo/middleware"
  "github.com/labstack/echo"
  "./handler"
)

func main() {
  e := echo.New()
  e.Use(middleware.Logger())  // ログ
  e.Use(middleware.Recover())
  e.Use(middleware.CORS()) // CORS対策
  initRouting(e) // ルーティング
}

func initRouting(e *echo.Echo) {
  e.GET("/questions", handler.GetAllQuestions)
  e.GET("/question/:id", handler.GetQuestion)
  e.POST("/question", handler.CreateQuestion)
  e.PUT("/question/:id", handler.UpdateQuestion)
  e.DELETE("/question/:id", handler.DeleteQuestion)

  e.GET("/users", handler.GetAllUsers)
  e.GET("/user/:id", handler.GetUser)
  e.POST("/user", handler.CreateUser)
  e.PUT("/user/:id", handler.UpdateUser)
  e.DELETE("/user/:id", handler.DeleteUser)

  e.GET("/comments", handler.GetAllComments)
  e.GET("/comment/:id", handler.GetComment)  // :idはCommentのIDではなくQuestionのID
  e.POST("/comment", handler.CreateComment)
  e.PUT("/comment/:id", handler.UpdateComment)
  e.DELETE("/comment/:id", handler.DeleteComment)

  e.GET("/mylist/:id", handler.GetMylistForUser)  // :idはMylistのIDではなくUserのID
  e.POST("/mylist", handler.CreateMylist)
  e.DELETE("/mylist/:id", handler.DeleteMylist)

  e.GET("/answer/:id", handler.GetAnswer)  // :idはAnswerのIDではなくQuestionのID
  e.POST("/answer", handler.CreateAnswer)
  e.PUT("/answer/:id", handler.UpdateAnswer)
  e.DELETE("/answer/:id", handler.DeleteAnswer)

  e.POST("/questionimage", handler.CreateQuestionImage)
  e.DELETE("/questionimage/:id", handler.DeleteQuestionImage)

  e.POST("/answerimage", handler.CreateAnswerImage)
  e.DELETE("/answerimage/:id", handler.DeleteAnswerImage)

  // e.POST("/association/:id", handler.CreateAssociation)
  // e.DELETE("/association/:id", handler.DeleteAssociation)

  e.POST("/tag", handler.CreateTag)

  e.Start(":1323")
}