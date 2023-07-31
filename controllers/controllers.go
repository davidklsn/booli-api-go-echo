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
	data := map[string]interface{}{
		"Users": users,
	}

	return c.Render(http.StatusOK, "start", data)
}

func User(c echo.Context) error {
	parseId := c.Param("id")
	user, err := handlers.GetUser(parseId)

	userJSON, err := json.Marshal(user)
	if err != nil {
		c.String(http.StatusNotFound, "Could not find user")
		return err
	}

	data := map[string]any{
		"UserID": user.UserID,
		"UserData": string(userJSON),
	}

	return c.Render(http.StatusOK, "user", data)
}

func ApiDocs(c echo.Context) error {
	return c.Render(http.StatusOK, "docs", nil)
}
