package api

import (
	"net/http"

	"github.com/davidklsn/booli-api-go/handlers"
	"github.com/davidklsn/booli-api-go/types"
	"github.com/labstack/echo/v4"
)

func HandleGetUsers(c echo.Context) error {
	users, err := handlers.GetUsers()

	if err != nil {
		c.String(http.StatusNotFound, "Could not find any users")
		return err
	}

	return c.JSON(http.StatusOK, users)
}

func HandleGetUser(c echo.Context) error {
	parseId := c.Param("id")
	user, err := handlers.GetUser(parseId)

	if err != nil {
		c.String(http.StatusNotFound, "Could not find user")
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func HandleCreateUser(c echo.Context) error {
	parseId := c.Param("id")
	req := new(types.Request)

	if err := c.Bind(req); err != nil {
		return err
	}

	userData, err := handlers.CreateUser(parseId, req.Residence, req.Activity, req.Info)

	if err != nil {
		c.Error(err)
		return err
	}

	return c.JSON(http.StatusOK, userData)
}

func HandleDeleteUser(c echo.Context) error {
	parseId := c.Param("id")
	req := new(types.Request)

	if err := c.Bind(req); err != nil {
		return err
	}

	userData, err := handlers.DeleteUser(parseId)

	if err != nil {
		c.Error(err)
		return err
	}

	return c.JSON(http.StatusOK, userData)
}

func HandleUpdateUserResidences(c echo.Context) error {
	parseId := c.Param("id")
	req := new(types.Request)

	if err := c.Bind(req); err != nil {
		return err
	}

	user, err := handlers.UpdateResidences(parseId, req.Residence)

	if err != nil {
		c.Error(err)
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func HandleUpdateUserActivities(c echo.Context) error {
	parseId := c.Param("id")
	req := new(types.Request)

	if err := c.Bind(req); err != nil {
		return err
	}

	user, err := handlers.UpdateActivites(parseId, req.Activity)

	if err != nil {
		c.Error(err)
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func HandleUpdateUserInfo(c echo.Context) error {
	parseId := c.Param("id")
	req := new(types.InfoRequest)

	if err := c.Bind(req); err != nil {
		return err
	}

	user, err := handlers.UpdateInfo(parseId, req.Info)

	if err != nil {
		c.Error(err)
		return err
	}

	return c.JSON(http.StatusOK, user)
}
