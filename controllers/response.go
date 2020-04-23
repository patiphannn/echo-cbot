package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/polnoy/echo-cbot/models"
	"github.com/polnoy/echo-cbot/services"
	"gopkg.in/mgo.v2/bson"
)

// Response is response service
type Response struct {
}

// Answer handler answer response.
func (h *Response) Answer(c echo.Context) (err error) {
	service := new(services.Response)
	profile := services.GetProfile(c)
	if profile["type"].(string) != "project" {
		return c.JSON(http.StatusBadRequest, bson.M{"message": "project key invalid!"})
	}

	form := &models.Response{}
	// skip checking bind errors.
	if err = c.Bind(form); err != nil {
		return c.JSON(http.StatusInternalServerError, bson.M{"message": err.Error()})
	}

	// Validate our data:
	if err = c.Validate(form); err != nil {
		return c.JSON(http.StatusInternalServerError, bson.M{"message": err.Error()})
	}

	result, err := service.Answer(form, profile)
	if err != nil {
		return c.JSON(http.StatusBadRequest, bson.M{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, result)
}
