package handler

import (
    "net/http"
    "github.com/labstack/echo"
    "github.com/jinzhu/gorm"
    "time"
    . "../controller"
)

type Model struct {
  ID        uint `gorm:"primary_key"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt *time.Time
}


type Question struct {
    gorm.Model
    Text string 
    UserID int   // 元はCreatorId  // ;gorm:"foreignkey:ID"
    User User  `gorm:"foreignkey:UserID"` // `gorm:"ForeignKey:UserID;AssociationForeignKey:ID"`
}


type User struct {
    gorm.Model
    Name string   // `json:"name"`
    Email string  // `json:"email"`
    Password string   // `json:"password"`
    Question []Question  // []Questionとすると動いた！！
}

type Tag struct {
    gorm.Model
    Name string
}


func GetAllQuestions(c echo.Context) error {
    db := OpenSQLiteConnection()
    defer db.Close()
    db.AutoMigrate(&Question{})

    var questions []Question
    db.Find(&questions)
    return c.JSON(http.StatusOK, questions)
}

func GetAllUsers(c echo.Context) error {
    db := OpenSQLiteConnection()
    defer db.Close()
    db.AutoMigrate(&User{})

    var users []User
    db.Find(&users)
    return c.JSON(http.StatusOK, users)
}

func GetQuestion(c echo.Context) error {
    db := OpenSQLiteConnection()
    defer db.Close()
    db.AutoMigrate(&Question{})

    if id := c.Param("id"); id != "" {
        var question Question
        db.First(&question, id).Related(&question.User)
        return c.JSON(http.StatusOK, question) 
    } else {
        return c.JSON(http.StatusNotFound, nil)
    }
}

func CreateQuestion(c echo.Context) error {
    db := OpenSQLiteConnection()
    defer db.Close()
    db.AutoMigrate(&Question{})

    question := new(Question)
    // question := new(Question{text:"さああ", creator_d: 1})
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