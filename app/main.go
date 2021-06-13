package main

import (
    "net/http"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "app/handler"
    "app/models/users"
    "app/models/user"
)

func main() {
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
        return c.String(http.StatusOK, "Hello galaxy!")
    })
    e.Logger.Fatal(e.Start(":8080"))
}
