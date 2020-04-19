package services

import (
	"time"

	"github.com/Kamva/mgm/v2"
	"github.com/dgrijalva/jwt-go"
	"github.com/polnoy/echo-cbot/models"
	"gopkg.in/mgo.v2/bson"
)

// Project is all project services
type Project struct{}

func createKey(project *models.Project) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["_id"] = project.ID.Hex()
	claims["name"] = project.Name
	claims["type"] = "project"
	claims["exp"] = time.Now().Add(0).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return t, nil
}

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
func (h *Project) Create(form *models.Project, profile jwt.MapClaims) (*models.Project, error) {
	key, err := createKey(form)
	if err != nil {
		return nil, err
	}

	userID := profile["_id"].(string)

	form.Key = key
	form.User = userID
	if err := mgm.Coll(form).Create(form); err != nil {
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
