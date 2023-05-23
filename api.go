package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func handleGetUsers(c echo.Context) error {
	users, err := GetUsers()

	if err != nil {
		c.String(http.StatusNotFound, "Could not find any users")
		return err
	}

	return c.JSON(http.StatusOK, users)
}

func handleGetUser(c echo.Context) error {
	parseId := c.Param("id")
	user, err := GetUser(parseId)

	if err != nil {
		c.String(http.StatusNotFound, "Could not find user")
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func handleUpdateUser(c echo.Context) error {
	parseId := c.Param("id")
	req := new(Request)

	if err := c.Bind(req); err != nil {
		return err
	}

	user, err := UpdateUser(parseId, req.ResidenceID, req.Meta)

	if err != nil {
		c.Error(err)
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func handleCreateUser(c echo.Context) error {
	parseId := c.Param("id")
	req := new(Request)

	if err := c.Bind(req); err != nil {
		return err
	}

	userData, err := CreateUser(parseId, req.ResidenceID, req.Meta)

	if err != nil {
		c.Error(err)
		return err
	}

	return c.JSON(http.StatusOK, userData)
}

func handleDeleteUser(c echo.Context) error {
	parseId := c.Param("id")
	req := new(Request)

	if err := c.Bind(req); err != nil {
		return err
	}

	userData, err := DeleteUser(parseId)

	if err != nil {
		c.Error(err)
		return err
	}

	return c.JSON(http.StatusOK, userData)
}
