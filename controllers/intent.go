package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/polnoy/echo-cbot/models"
	"github.com/polnoy/echo-cbot/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

// Intent is intent service
type Intent struct {
}

// Gets defined find all intent.
func (h *Intent) Gets(c echo.Context) (err error) {
	service := new(services.Intent)
	profile := services.GetProfile(c)
	projID, _ := primitive.ObjectIDFromHex(profile["_id"].(string))

	if profile["type"].(string) != "project" {
		return c.JSON(http.StatusBadRequest, bson.M{"message": "project key invalid!"})
	}

	result, err := service.Gets(bson.M{"project": projID})
	if err != nil {
		return c.JSON(http.StatusBadRequest, bson.M{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

// Get defined find once intent.
func (h *Intent) Get(c echo.Context) (err error) {
	service := new(services.Intent)
	profile := services.GetProfile(c)
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	projID, _ := primitive.ObjectIDFromHex(profile["_id"].(string))

	if profile["type"].(string) != "project" {
		return c.JSON(http.StatusBadRequest, bson.M{"message": "project key invalid!"})
	}

	result, err := service.Get(bson.M{"_id": id, "project": projID})
	if err != nil {
		return c.JSON(http.StatusBadRequest, bson.M{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

// Create handler create new intent.
func (h *Intent) Create(c echo.Context) (err error) {
	service := new(services.Intent)
	profile := services.GetProfile(c)

	if profile["type"].(string) != "project" {
		return c.JSON(http.StatusBadRequest, bson.M{"message": "project key invalid!"})
	}

	form := &models.Intent{}
	// skip checking bind errors.
	if err = c.Bind(form); err != nil {
		return c.JSON(http.StatusInternalServerError, bson.M{"message": err.Error()})
	}

	// Validate our data:
	if err = c.Validate(form); err != nil {
		return c.JSON(http.StatusInternalServerError, bson.M{"message": err.Error()})
	}

	result, err := service.Create(form, profile)
	if err != nil {
		return c.JSON(http.StatusBadRequest, bson.M{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, result)
}

// Update handler update intent.
func (h *Intent) Update(c echo.Context) (err error) {
	service := new(services.Intent)
	profile := services.GetProfile(c)
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	projID, _ := primitive.ObjectIDFromHex(profile["_id"].(string))

	if profile["type"].(string) != "project" {
		return c.JSON(http.StatusBadRequest, bson.M{"message": "project key invalid!"})
	}

	form := &models.RIntent{}
	// skip checking bind errors.
	if err = c.Bind(form); err != nil {
		return c.JSON(http.StatusInternalServerError, bson.M{"message": err.Error()})
	}

	// Validate our data:
	if err = c.Validate(form); err != nil {
		return c.JSON(http.StatusInternalServerError, bson.M{"message": err.Error()})
	}

	result, err := service.Update(bson.M{"_id": id, "project": projID}, form)
	if err != nil {
		return c.JSON(http.StatusBadRequest, bson.M{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

// Delete handler delete intent.
func (h *Intent) Delete(c echo.Context) error {
	service := new(services.Intent)
	profile := services.GetProfile(c)
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	projID, _ := primitive.ObjectIDFromHex(profile["_id"].(string))

	if profile["type"].(string) != "project" {
		return c.JSON(http.StatusBadRequest, bson.M{"message": "project key invalid!"})
	}

	if err := service.Delete(bson.M{"_id": id, "project": projID}); err != nil {
		return c.JSON(http.StatusBadRequest, bson.M{"message": err.Error()})
	}

	return c.JSON(http.StatusNoContent, nil)
}
