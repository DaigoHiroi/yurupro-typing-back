package models/users

import (
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

type (
	User struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	Response struct {
		Users []User `json:"users"`
	}
)

var (
	dsn = "user:password@tcp(127.0.0.1:3306)/echo?charset=utf8mb4&parseTime=True&loc=Local"
)

func SelectUsers(c echo.Context) error {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	var users []User

	db.Find(&users)

	response := new(Response)
	response.Users = users

	return c.JSON(http.StatusOK, response)
}
