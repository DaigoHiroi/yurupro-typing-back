package main

import (
    "fmt"
    "net/http"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "app/handler"
    "app/models/users"
    "app/models/user"
    "time"
    "github.com/jinzhu/gorm"
)

func main() {
    db := sqlConnect()
    defer db.Close()

    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.POST("/login", handler.Login())
    r := e.Group("/restricted")
    r.Use(middleware.JWT([]byte("secret")))
    r.GET("/welcome", handler.Restricted())

    e.GET("/users/:id", user.SelectUser)
    e.GET("/users", users.SelectUsers)

    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello galaxyff")
    })
    e.Logger.Fatal(e.Start(":8080"))
}

func sqlConnect() (database *gorm.DB) {
    DBMS := "mysql"
    USER := "echo"
    PASS := "golang"
    PROTOCOL := "tcp(mysql:3306)"
    DBNAME := "echo"

    CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

    count := 0
    db, err := gorm.Open(DBMS, CONNECT)
    if err != nil {
        for {
            if err == nil {
                fmt.Println("")
                break
            }
            fmt.Print(".")
            time.Sleep(time.Second)
            count++
            if count > 10 {
                fmt.Println("")
                fmt.Println("DB接続失敗")
                fmt.Println(err.Error())
                panic(err)
            }
            db, err = gorm.Open(DBMS, CONNECT)
        }
    }
    fmt.Println("DB接続成功")

    return db
}
