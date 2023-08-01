package controllers

import (
	"encoding/json"
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

func User(c echo.Context) error {
	parseId := c.Param("id")
	user, err := handlers.GetUser(parseId)

	userJSON, err := json.Marshal(user)
	if err != nil {
		c.String(http.StatusNotFound, "Could not find user")
		return err
	}

	return c.Render(http.StatusOK, "user", echo.Map{
		"title":    "Användare",
		"userID":   user.UserID,
		"userData": string(userJSON),
	})
}

func ApiDocs(c echo.Context) error {
	return c.Render(http.StatusOK, "docs", echo.Map{
		"title": "API Docs",
	})
}
