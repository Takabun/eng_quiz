package handler

// import (
// 	"net/http"
// 	"github.com/labstack/echo"
// 	. "../controller"
// 	. "../model"
// )

// func CreateAssociation(c echo.Context) error {
//   db := OpenSQLiteConnection()
//   defer db.Close()
//   // db.AutoMigrate(&Mylist{})

//   mylist := new(Mylist)
//   if err := c.Bind(mylist); err != nil {
//     return err
//   }
//   db.Create(&mylist)

//   return c.JSON(http.StatusOK, mylist)
// }

// func DeleteAssociation(c echo.Context) error {
//   db := OpenSQLiteConnection()
//   defer db.Close()

//   if id := c.Param("id"); id != "" {
//     var mylist Mylist
//     db.First(&mylist, id)
//     db.Delete(mylist)
//     return c.JSON(http.StatusOK, mylist)
//   } else {
//     return c.JSON(http.StatusNotFound, nil)
//   }
// }

// // func CreateAssociation(c echo.Context) error {
// //   db := OpenSQLiteConnection()
// //   defer db.Close()
// //   db.AutoMigrate(&Mylist{})

// //   mylist := new(Mylist)
// //   if err := c.Bind(mylist); err != nil {
// //     return err
// //   }
// //   db.Create(&mylist)

// //   return c.JSON(http.StatusOK, mylist)
// // }

// // func DeleteAssociation(c echo.Context) error {
// //   db := OpenSQLiteConnection()
// //   defer db.Close()

// //   if id := c.Param("id"); id != "" {
// //     var mylist Mylist
// //     db.First(&mylist, id)
// //     db.Delete(mylist)
// //     return c.JSON(http.StatusOK, mylist)
// //   } else {
// //     return c.JSON(http.StatusNotFound, nil)
// //   }
// // }