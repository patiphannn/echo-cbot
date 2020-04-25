package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/polnoy/echo-cbot/models"
	"github.com/polnoy/echo-cbot/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

// Project is project service
type Project struct {
}

// Gets defined find all project.
func (h *Project) Gets(c echo.Context) (err error) {
	service := new(services.Project)
	profile := services.GetProfile(c)
	user, _ := primitive.ObjectIDFromHex(profile["_id"].(string))

	result, err := service.Gets(bson.M{"user": user})
	if err != nil {
		return c.JSON(http.StatusBadRequest, bson.M{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

// Get defined find once project.
func (h *Project) Get(c echo.Context) (err error) {
	service := new(services.Project)
	profile := services.GetProfile(c)
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	user, _ := primitive.ObjectIDFromHex(profile["_id"].(string))

	result, err := service.Get(bson.M{"_id": id, "user": user})
	if err != nil {
		return c.JSON(http.StatusBadRequest, bson.M{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

// Create handler create new project.
func (h *Project) Create(c echo.Context) (err error) {
	service := new(services.Project)
	profile := services.GetProfile(c)
	user, _ := primitive.ObjectIDFromHex(profile["_id"].(string))

	form := &models.Project{}
	// skip checking bind errors.
	if err = c.Bind(form); err != nil {
		return c.JSON(http.StatusInternalServerError, bson.M{"message": err.Error()})
	}

	// Validate our data:
	if err = c.Validate(form); err != nil {
		return c.JSON(http.StatusInternalServerError, bson.M{"message": err.Error()})
	}

	form.User = user
	result, err := service.Create(form)
	if err != nil {
		return c.JSON(http.StatusBadRequest, bson.M{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, result)
}

// Update handler update project.
func (h *Project) Update(c echo.Context) (err error) {
	service := new(services.Project)
	profile := services.GetProfile(c)
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	user, _ := primitive.ObjectIDFromHex(profile["_id"].(string))

	form := &models.RProject{}
	// skip checking bind errors.
	if err = c.Bind(form); err != nil {
		return c.JSON(http.StatusInternalServerError, bson.M{"message": err.Error()})
	}

	// Validate our data:
	if err = c.Validate(form); err != nil {
		return c.JSON(http.StatusInternalServerError, bson.M{"message": err.Error()})
	}

	result, err := service.Update(bson.M{"_id": id, "user": user}, form)
	if err != nil {
		return c.JSON(http.StatusBadRequest, bson.M{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

// Delete handler delete project.
func (h *Project) Delete(c echo.Context) error {
	service := new(services.Project)
	profile := services.GetProfile(c)
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	user, _ := primitive.ObjectIDFromHex(profile["_id"].(string))

	if err := service.Delete(bson.M{"_id": id, "user": user}); err != nil {
		return c.JSON(http.StatusBadRequest, bson.M{"message": err.Error()})
	}

	return c.JSON(http.StatusNoContent, nil)
}
