package controller

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

func OpenSQLiteConnection() *gorm.DB {
    // db, err := gorm.Open("mysql",  "root:@/eng_quiz?parseTime=true&loc=Asia%2FTokyo")
    db, err := gorm.Open("mysql", "root:password@tcp(db:3306)/eng_quiz?parseTime=true&loc=Asia%2FTokyo")
    if err != nil {
        panic("failed to connect database.")
    }
    db.LogMode(true)
    return db
}