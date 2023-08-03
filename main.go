package main

import (
	"fmt"
	"os"

	"github.com/davidklsn/booli-api-go/api"
	"github.com/davidklsn/booli-api-go/config"
	"github.com/davidklsn/booli-api-go/constants"
	"github.com/davidklsn/booli-api-go/controllers"
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := config.LoadENV()
	if err != nil {
		panic(err)
	}

	// Initalize database
	constants.InitDB()

	e := echo.New()
	e.Renderer = echoview.Default()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/dist", "public/dist")

	// -- API --
	// Routes [:users]
	g := e.Group("/users")
	e.GET("/users", api.HandleGetUsers)
	g.GET("/:id", api.HandleGetUser)
	g.POST("/:id", api.HandleCreateUser)
	g.DELETE("/:id", api.HandleDeleteUser)


	// Routes [:custom_data]
	g.PUT("/:id/update_residences", api.HandleUpdateUserResidences)
	g.PUT("/:id/update_info", api.HandleUpdateUserInfo)

	g.GET("/:id/current_residence", api.HandleGetCurrentUserResidence)

	// -- PAGES --
	e.GET("/", controllers.Index)
	e.GET("/u/:id", controllers.User)
	e.GET("/docs", controllers.ApiDocs)

	// partial routes
	e.GET("/users/:id/edit", controllers.EditUser)

	port := os.Getenv("APP_PORT")
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s%s", ":", port)))
}
