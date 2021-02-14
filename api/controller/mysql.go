package controller

import (
    "os"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

func OpenSQLiteConnection() *gorm.DB {
    dbUrl := os.Getenv("DB_URL")
    if dbUrl == "" {
      dbUrl = "root:password@tcp(db:3306)/eng_quiz?parseTime=true&loc=Asia%2FTokyo"
	}
    
    db, err := gorm.Open("mysql", dbUrl)
    if err != nil {
        panic("failed to connect database.")
    }
    db.LogMode(true)
    return db
}