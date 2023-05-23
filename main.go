package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/davidklsn/booli-api-go/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := config.LoadENV()
	if err != nil {
		panic(err)
	}

	InitDB()
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	g := e.Group("/user")
	g.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		return key == os.Getenv("AUTH_KEY"), nil
	}))


	g.GET("/:id", handleGetUser)
	g.POST("/:id", handleCreateUser)
	g.PUT("/:id", handleUpdateUser)
	g.DELETE("/:id", handleDeleteUser)

	e.GET("/users", handleGetUsers)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "A simple API to save and retrieve user data. Endpoint is /user/:id")
	})

	port := os.Getenv("APP_PORT")
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s%s", ":", port)))
}