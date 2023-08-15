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
	e.Static("/img", "public/img")


	// -- API --
	// Routes [:users]
	g := e.Group("/users")

	e.GET("/users", api.HandleGetUsers)
	e.GET("/users/search", api.HandleSearchUsers)
	g.GET("/:id", api.HandleGetUser)
	g.POST("/:id", api.HandleCreateUser)
	g.DELETE("/:id", api.HandleDeleteUser)

	// Routes [:custom_data]
	g.PUT("/:id/update_residences", api.HandleUpdateUserResidences)
	g.PUT("/:id/update_current_residence", api.HandleUpdateUserCurrentResidence)
	g.PUT("/:id/update_selected_residence", api.HandleUpdateUserSelectedResidence)
	g.PUT("/:id/update_info", api.HandleUpdateUserInfo)

	g.DELETE("/:id/delete_residence", api.HandleDeleteUserResidence)

	g.GET("/:id/current_residence", api.HandleGetCurrentUserResidence)
	g.GET("/:id/selected_residence", api.HandleGetSelectedUserResidence)

	// -- PAGES --
	e.GET("/", controllers.Index)
	e.GET("/docs", controllers.ApiDocs)

	//Users
	e.GET("/u/:id", controllers.User)
	e.GET("/u/search", controllers.SearchUsers)
	e.GET("/u/:id/edit", controllers.EditUser)
	e.POST("/u/:id/edit", controllers.EditUser)

	port := os.Getenv("APP_PORT")
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s%s", ":", port)))
}
