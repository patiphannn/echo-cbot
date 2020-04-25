package common

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/polnoy/echo-cbot/models"
	"github.com/spf13/viper"
)

// CreateKey defind create project key
func CreateKey(project *models.Project) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["_id"] = project.ID.Hex()
	claims["name"] = project.Name
	claims["type"] = "project"

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(viper.GetString("access_key")))
	if err != nil {
		return "", err
	}

	return t, nil
}
