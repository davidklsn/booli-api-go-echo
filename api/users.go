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

func HandleSearchUsers(c echo.Context) error {
	ids := c.QueryParam("id")
	users, err := handlers.GetUsersByIds(ids)

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

	userData, err := handlers.CreateUser(parseId, req.Residence, req.Info)

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

func HandleDeleteUserResidence(c echo.Context) error {
	parseId := c.Param("id")
	req := new(types.Request)

	if err := c.Bind(req); err != nil {
		return err
	}

	user, err := handlers.DeleteUserResidence(parseId, req.Residence)

	if err != nil {
		c.Error(err)
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func HandleUpdateUserCurrentResidence(c echo.Context) error {
	parseId := c.Param("id")
	req := new(types.Request)

	if err := c.Bind(req); err != nil {
		return err
	}

	user, err := handlers.UpdateCurrentResidence(parseId, req.Residence)

	if err != nil {
		c.Error(err)
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func HandleUpdateUserSelectedResidence(c echo.Context) error {
	parseId := c.Param("id")
	req := new(types.Request)

	if err := c.Bind(req); err != nil {
		return err
	}

	user, err := handlers.UpdateSelectedResidence(parseId, req.Residence)

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

func HandleGetCurrentUserResidence(c echo.Context) error {
	parseId := c.Param("id")

	residence, err := handlers.GetCurrentResidence(parseId)

	if err != nil {
		return c.JSON(http.StatusOK, map[string]any{})
	}

	return c.JSON(http.StatusOK, residence)
}

func HandleGetSelectedUserResidence(c echo.Context) error {
	parseId := c.Param("id")

	residence, err := handlers.GetSelectedResidence(parseId)

	if err != nil {
		return c.JSON(http.StatusOK, map[string]any{})
	}

	return c.JSON(http.StatusOK, residence)
}
