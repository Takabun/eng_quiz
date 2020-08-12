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
  CreatedAt *time.Time
  UpdatedAt *time.Time
  DeletedAt *time.Time
}


type Question struct {
    gorm.Model
    Text string 
    UserID int
    User User  // `gorm:"foreignkey:UserID"`なくてもイケる
    Tags []*Tag  `gorm:"many2many:question_tags;"`
    Answer Answer
    Comments []Comment
}

type Answer struct {
    gorm.Model
    Text string
    QuestionID int
    // Question Question  //不要。(指定のQuestionのページにしか現れないと分かっているので)
}

type User struct {
    gorm.Model
    Name string
    Email string
    Password string
    Questions []Question  //[]Questionとすると(スライス)動く
}

type Tag struct {
    gorm.Model
    Name string
    Questions []*Question  `gorm:"many2many:question_tags;"`
}

type Comment struct {
    gorm.Model
    Text string
    QuestionID int //実際はたぶん使わない
    UserID int
    User User // 動作しない`gorm:"foreignkey:UserID"`
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
        db.First(&question, id).Related(&question.User).Related(&question.Answer).Related(&question.Comments).Related(&question.Tags, "Tags")
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


//  Users
func GetAllUsers(c echo.Context) error {
    db := OpenSQLiteConnection()
    defer db.Close()
    db.AutoMigrate(&User{})

    var users []User
    db.Find(&users)
    return c.JSON(http.StatusOK, users)
}

// 個別のAnswer(:idはAnswerのIDではなくQuestionのID)
func GetAnswer(c echo.Context) error {
    db := OpenSQLiteConnection()
    defer db.Close()
    db.AutoMigrate(&Answer{})

    if id := c.Param("id"); id != "" {
        var answer Answer
        db.Where("question_id = ?", id).First(&answer)
        return c.JSON(http.StatusOK, answer)
    } else {
        return c.JSON(http.StatusNotFound, nil)
    }
}


// 個別のComment(:idはCommentのIDではなくQuestionのID)
func GetComment(c echo.Context) error {
    db := OpenSQLiteConnection()
    defer db.Close()
    db.AutoMigrate(&Comment{})

    if id := c.Param("id"); id != "" {
        var comments []Comment      // Commentは複数あるので[]
        db.Find(&comments).Where("question_id = ?", id)

        for i := range comments {
            db.Model(comments[i]).Related(&comments[i].User) // 第二引数("Users")なくてもイケた
        }

        return c.JSON(http.StatusOK, comments)
    } else {
        return c.JSON(http.StatusNotFound, nil)
    }
}
