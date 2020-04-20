package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/polnoy/echo-cbot/models"
	"github.com/polnoy/echo-cbot/services"
	"gopkg.in/mgo.v2/bson"
)

// Auth is auth service
type Auth struct {
}

// Signup handler create new user.
func (h *Auth) Signup(c echo.Context) (err error) {
	service := new(services.Auth)
	form := &models.Signup{}
	// skip checking bind errors.
	if err := c.Bind(form); err != nil {
		return c.JSON(http.StatusInternalServerError, bson.M{"message": err.Error()})
	}

	// Validate our data:
	if err := c.Validate(form); err != nil {
		return c.JSON(http.StatusInternalServerError, bson.M{"message": err.Error()})
	}

	data, err := service.Signup(form)
	if err != nil {
		return c.JSON(http.StatusBadRequest, bson.M{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}

// Signin handler sigin user.
func (h *Auth) Signin(c echo.Context) (err error) {
	service := new(services.Auth)
	form := &models.Signin{}
	// skip checking bind errors.
	if err := c.Bind(form); err != nil {
		return c.JSON(http.StatusInternalServerError, bson.M{"message": err.Error()})
	}

	// Validate our data:
	if err := c.Validate(form); err != nil {
		return c.JSON(http.StatusInternalServerError, bson.M{"message": err.Error()})
	}

	data, err := service.Signin(form)
	if err != nil {
		return c.JSON(http.StatusBadRequest, bson.M{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}
