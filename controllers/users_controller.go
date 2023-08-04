package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/davidklsn/booli-api-go/handlers"
	"github.com/foolin/goview"
	"github.com/labstack/echo/v4"
)

func User(c echo.Context) error {
	parseId := c.Param("id")
	user, err := handlers.GetUser(parseId)

	userJSON, err := json.Marshal(user)
	if err != nil {
		c.String(http.StatusNotFound, "Could not find user")
		return err
	}

	var view string
	if c.Request().Header.Get("HX-Request") == "true" {
		view = "/users/partials/show"
	} else {
		view = "/users/show"
	}

	return c.Render(http.StatusOK, view, goview.M{
		"title":    "Användare",
		"userID":   user.UserID,
		"userData": string(userJSON),
	})
}

func EditUser(c echo.Context) error {
	parseId := c.Param("id")
	user, err := handlers.GetUser(parseId)
	userJSON, err := json.Marshal(user)

	if err != nil {
		c.String(http.StatusNotFound, "Could not find user")
		return err
	}

	var view string
	if c.Request().Header.Get("HX-Request") == "true" {
		view = "/users/partials/edit"
	} else {
		view = "/users/edit"
	}

	return c.Render(http.StatusOK, view, goview.M{
		"title":    "Redigera användare",
		"userID":   user.UserID,
		"userData": string(userJSON),
	})
}
