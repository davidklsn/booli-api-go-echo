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
	data := map[string]interface{}{
		"Users": users,
	}

	return c.Render(http.StatusOK, "start", data)
}

func ApiExplorer(c echo.Context) error {
	return c.Render(http.StatusOK, "explorer", nil)
}
