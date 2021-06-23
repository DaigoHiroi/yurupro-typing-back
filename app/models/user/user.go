package user

import (
	//"github.com/labstack/echo/v4"
	"github.com/labstack/echo"
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
		User User `json:"user"`
	}
)

var (
	dsn = "echo:golang@tcp(mysql:3306)/echo?charset=utf8mb4&parseTime=True&loc=Local"
)

func SelectUser(c echo.Context) error {
	id := c.Param("id")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	var user User

	db.First(&user, id)

	response := new(Response)
	response.User = user

	return c.JSON(http.StatusOK, response)
}
