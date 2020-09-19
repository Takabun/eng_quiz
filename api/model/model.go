package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Model struct {
  ID        uint `gorm:"primary_key"`
  CreatedAt *time.Time
  UpdatedAt *time.Time
  DeletedAt *time.Time
}

type Question struct {
	gorm.Model
	User string
	Text string 
	DefaultImage int
	// UserID int
	// User User  // `gorm:"foreignkey:UserID"`なくてもイケる
	Tags []Tag  `gorm:"many2many:question_tags;"`
	QuestionImages []QuestionImage
	// Answer Answer
	// Comments []Comment	
}

type QuestionImage struct {
	gorm.Model
	Url string
	QuestionID int
	Name string
}

type Answer struct {
	gorm.Model
	Text string
	QuestionID int
	// Question Question  //不要。(指定のQuestionのページにしか現れないと分かっているので)
	AnswerImage []AnswerImage
}

type AnswerImage struct {
	gorm.Model
	Url string
	AnswerID int
	Name string
}

// type User struct {
// 	gorm.Model
// 	Name string
// 	Email string
// 	Password string
// 	Questions []Question  //[]Questionとすると(スライス)動く
// }

type Tag struct {
	gorm.Model
	Name string
	Questions []*Question  `gorm:"many2many:question_tags;"`
}

type Comment struct {
	gorm.Model
	User string
	Text string
	QuestionID int //実際はたぶん使わない
	// UserID int
	// User User // 動作しない`gorm:"foreignkey:UserID"`
	
}

// type Mylist struct {
// 	gorm.Model
// 	UserID int
// 	QuestionID int
// 	User User
// 	Question Question
// }
