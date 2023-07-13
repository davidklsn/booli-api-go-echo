package api

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

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

func HandleUpdateUser(c echo.Context) error {
	parseId := c.Param("id")
	req := new(types.Request)

	if err := c.Bind(req); err != nil {
		return err
	}

	user, err := handlers.UpdateUser(parseId, req.Residence, req.Activity, req.Info)

	if err != nil {
		c.Error(err)
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

func HandleUpdateResidences(c echo.Context) error {
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

func TestHandleUpdateResidences(t *testing.T) {
	// Create a new echo context for testing
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/residences/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Set up the request body
	reqBody := `{"residence": "New Residence"}`
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Body = ioutil.NopCloser(strings.NewReader(reqBody))

	// Call the handler function
	err := HandleUpdateResidences(c)

	// Check if there was an error
	if err != nil {
		t.Errorf("handleUpdateResidences returned an error: %v", err)
	}

	// Check the response status code
	if rec.Code != http.StatusOK {
		t.Errorf("handleUpdateResidences returned wrong status code: got %v, want %v", rec.Code, http.StatusOK)
	}

	// Check the response body
	expectedResponseBody := `{"id": "1", "residence": "New Residence"}`
	if rec.Body.String() != expectedResponseBody {
		t.Errorf("handleUpdateResidences returned wrong response body: got %v, want %v", rec.Body.String(), expectedResponseBody)
	}
}

func handleUpdateActivity(c echo.Context) error {
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
