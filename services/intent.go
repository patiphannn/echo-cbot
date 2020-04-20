package services

import (
	"fmt"

	"github.com/Kamva/mgm/v2"
	"github.com/dgrijalva/jwt-go"
	"github.com/polnoy/echo-cbot/models"
	"gopkg.in/mgo.v2/bson"
)

// Intent is all intent services
type Intent struct{}

// Gets defined find all intent.
func (h *Intent) Gets(cond bson.M) ([]models.Intent, error) {
	result := []models.Intent{}
	if err := mgm.Coll(&models.Intent{}).SimpleFind(&result, cond); err != nil {
		return nil, err
	}

	return result, nil
}

// Get defined find once intent.
func (h *Intent) Get(cond bson.M) (*models.Intent, error) {
	result := &models.Intent{}
	if err := mgm.Coll(result).First(cond, result); err != nil {
		return nil, err
	}

	return result, nil
}

// Create defined create new intent.
func (h *Intent) Create(form *models.Intent, profile jwt.MapClaims) (*models.Intent, error) {
	projectID := profile["_id"].(string)
	fmt.Println("projectID: ", projectID)

	form.Project = projectID
	if err := mgm.Coll(form).Create(form); err != nil {
		return nil, err
	}

	return form, nil
}

// Update defined update intent.
func (h *Intent) Update(cond bson.M, form *models.RIntent) (*models.Intent, error) {
	result := &models.Intent{}
	coll := mgm.Coll(result)
	if err := coll.First(cond, result); err != nil {
		return nil, err
	}
	result.Name = form.Name
	result.Context = form.Context
	result.Question = form.Question
	result.Answer = form.Answer
	result.Extra = form.Extra
	result.Fallback = form.Fallback
	result.Fulfillment = form.Fulfillment
	result.Welcome = form.Welcome

	if err := coll.Update(result); err != nil {
		return nil, err
	}
	return result, nil
}

// Delete defined delete intent.
func (h *Intent) Delete(cond bson.M) error {
	result := &models.Intent{}
	coll := mgm.Coll(result)
	if err := coll.First(cond, result); err != nil {
		return err
	}

	if err := coll.Delete(result); err != nil {
		return err
	}

	return nil
}
