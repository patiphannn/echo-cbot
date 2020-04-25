package services

import (
	"github.com/Kamva/mgm/v2"
	"github.com/polnoy/echo-cbot/common"
	"github.com/polnoy/echo-cbot/models"
	"gopkg.in/mgo.v2/bson"
)

// Project is all project services
type Project struct{}

// Gets defined find all project.
func (h *Project) Gets(cond bson.M) ([]models.Project, error) {
	result := []models.Project{}
	if err := mgm.Coll(&models.Project{}).SimpleFind(&result, cond); err != nil {
		return nil, err
	}

	return result, nil
}

// Get defined find once project.
func (h *Project) Get(cond bson.M) (*models.Project, error) {
	result := &models.Project{}
	if err := mgm.Coll(result).First(cond, result); err != nil {
		return nil, err
	}

	return result, nil
}

// Create defined create new project.
func (h *Project) Create(form *models.Project) (*models.Project, error) {
	coll := mgm.Coll(form)
	if err := coll.Create(form); err != nil {
		return nil, err
	}

	key, err := common.CreateKey(form)
	if err != nil {
		return nil, err
	}
	form.Key = key
	if err := coll.Update(form); err != nil {
		return nil, err
	}

	return form, nil
}

// Update defined update project.
func (h *Project) Update(cond bson.M, form *models.RProject) (*models.Project, error) {
	result := &models.Project{}
	coll := mgm.Coll(result)
	if err := coll.First(cond, result); err != nil {
		return nil, err
	}
	result.Name = form.Name
	result.Fulfillment = form.Fulfillment
	result.Enabled = form.Enabled

	if err := coll.Update(result); err != nil {
		return nil, err
	}
	return result, nil
}

// Delete defined delete project.
func (h *Project) Delete(cond bson.M) error {
	result := &models.Project{}
	coll := mgm.Coll(result)
	if err := coll.First(cond, result); err != nil {
		return err
	}

	if err := coll.Delete(result); err != nil {
		return err
	}

	return nil
}
