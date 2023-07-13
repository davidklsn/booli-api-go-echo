package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/davidklsn/booli-api-go/api"
	"github.com/davidklsn/booli-api-go/config"
	"github.com/davidklsn/booli-api-go/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := config.LoadENV()
	if err != nil {
		panic(err)
	}

	handlers.InitDB()
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	g := e.Group("/user")
	if os.Getenv("GO_ENV") == "production" {
		g.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
			return key == os.Getenv("AUTH_KEY"), nil
		}))
	}

	g.GET("/:id", api.HandleGetUser)
	g.POST("/:id", api.HandleCreateUser)
	g.PUT("/:id", api.HandleUpdateUser)
	g.DELETE("/:id", api.HandleDeleteUser)

	// Residences
	g.PUT("/:id/update_residence", api.HandleUpdateResidences)

	// 
	// g.PUT("/:id/update_activity", handleUpdateActivities)

	e.GET("/users", api.HandleGetUsers)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "A simple API to save and retrieve user data. Endpoint is /user/:id")
	})

	port := os.Getenv("APP_PORT")
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s%s", ":", port)))
}
