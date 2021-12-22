package main

import (
	"flag"
	"fmt"
	_ "github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"net/http"

	conf "app/conf"
	interactor "app/interactor"
	middleware "app/presenter/http/middleware"
	router "app/presenter/http/router"
)

//Dockerコンテナで実行する時(production.confをもとに起動するとき)は起動時に-serverを指定
var runServer = flag.Bool("server", false, "production is -server option require")

func main() {
	flag.Parse()
	conf.NewConfig(*runServer)

	// Echo instance
	e := echo.New()
	conn := conf.NewDBConnection()
	defer func() {
		if err := conn.Close(); err != nil {
			e.Logger.Fatal(fmt.Sprintf("Failed to close: %v", err))
		}
	}()
	i := interactor.NewInteractor(conn)
	h := i.NewAppHandler()

	router.NewRouter(e, h)
	middleware.NewMiddleware(e)
	if err := e.Start(fmt.Sprintf(":%d", conf.Current.Server.Port)); err != nil {
		e.Logger.Fatal(fmt.Sprintf("Failed to start: %v", err))
	}

	// Middleware
	// e.Use(middleware.Logger())
	//e.Use(middleware.Recover())

	// Routes
	//e.GET("/", hello)

	// Start server
	//e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World2!")
}
