package main

import (
	"fmt"
	"html/template"
	"io"
	"os"

	"github.com/davidklsn/booli-api-go/api"
	"github.com/davidklsn/booli-api-go/config"
	"github.com/davidklsn/booli-api-go/constants"
	"github.com/davidklsn/booli-api-go/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data any, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	err := config.LoadENV()
	if err != nil {
		panic(err)
	}

	// Initalize database
	constants.InitDB()

	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e := echo.New()
	e.Renderer = t

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	g := e.Group("/user")
	if os.Getenv("GO_ENV") == "production" {
		g.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
			return key == os.Getenv("AUTH_KEY"), nil
		}))
	}

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

	port := os.Getenv("APP_PORT")
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s%s", ":", port)))
}
