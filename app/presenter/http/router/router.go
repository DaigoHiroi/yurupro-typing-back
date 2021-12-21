package router

import (
	"app/presenter/http/handler"
	"github.com/labstack/echo/v4"
)

// NewRouter Routerの設定を行います.
func NewRouter(e *echo.Echo, h handler.AppHandler) {
	e.POST("/users", h.CreateUser)
	e.GET("/users", h.GetUsers)
	e.GET("/users/:id", h.GetUser)
	e.PUT("/users/:id", h.UpdateUser)
	e.DELETE("/users/:id", h.DeleteUser)
}
