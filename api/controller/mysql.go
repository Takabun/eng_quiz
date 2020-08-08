package controller

import (
    "github.com/jinzhu/gorm"
      _ "github.com/jinzhu/gorm/dialects/mysql"
)

func OpenSQLiteConnection() *gorm.DB {
    db, err := gorm.Open("mysql",  "root:@/go_trial?parseTime=true&loc=Asia%2FTokyo")
    if err != nil {
        panic("failed to connect database.")
    }
    db.LogMode(true)
    return db
}