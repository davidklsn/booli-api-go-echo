package controllers

import (
	"net/http"

	"github.com/davidklsn/booli-api-go/handlers"
	"github.com/davidklsn/booli-api-go/types"
	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	var users []types.UserData
	users, _ = handlers.GetUsers()

	return c.Render(http.StatusOK, "index", echo.Map{
		"title": "Start",
		"users": users,
	})
}

func ApiDocs(c echo.Context) error {
	return c.Render(http.StatusOK, "docs", echo.Map{
		"title": "API Docs",
	})
}
