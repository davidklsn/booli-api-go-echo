package main

import (
	"fmt"
	"net/http"
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

	g := e.Group("/users")
	// if os.Getenv("GO_ENV") == "production" {
	// 	g.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
	// 		return key == os.Getenv("AUTH_KEY"), nil
	// 	}))
	// }

	// Routes [:users]
	g.GET("/:id", api.HandleGetUser)
	g.POST("/:id", api.HandleCreateUser)
	g.DELETE("/:id", api.HandleDeleteUser)

	// Routes [:custom_data]
	g.PUT("/:id/update_residence", api.HandleUpdateUserResidences)
	g.PUT("/:id/update_activity", api.HandleUpdateUserActivities)
	g.PUT("/:id/update_info", api.HandleUpdateUserInfo)

	e.GET("/users", api.HandleGetUsers)
	e.GET("/", controllers.Index)
	e.GET("/u/:id", controllers.User)
	e.GET("/docs", controllers.ApiDocs)

	e.GET("/page", func(c echo.Context) error {
		//render only file, must full name with extension
		return c.Render(http.StatusOK, "page.html", echo.Map{"title": "Page file title!!"})
	})

	port := os.Getenv("APP_PORT")
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s%s", ":", port)))
}
