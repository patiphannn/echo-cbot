package services

import (
	"fmt"
	"time"

	"github.com/Kamva/mgm/v2"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/polnoy/echo-cbot/common"
	"github.com/polnoy/echo-cbot/models"
	"gopkg.in/mgo.v2/bson"
)

// Auth is all auth services
type Auth struct{}

func createToken(user *models.User) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["_id"] = user.ID.Hex()
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["admin"] = user.Admin
	claims["type"] = "user"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return t, nil
}

// GetProfile handler get user profile.
func GetProfile(c echo.Context) jwt.MapClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	return claims
}

// Signup defined signup user
func (h *Auth) Signup(form *models.Signup) (*models.User, error) {
	user := &models.User{}
	if err := mgm.Coll(user).First(bson.M{"email": form.Email}, user); err != nil && err.Error() != "mongo: no documents in result" {
		return nil, err
	}

	if user.Email != "" {
		return nil, fmt.Errorf("Email already exists: %s", user.Email)
	}

	password, err := common.GeneratePasswordHash([]byte(form.Password))
	if err != nil {
		return nil, err
	}

	data := &models.User{
		Email:    form.Email,
		Password: password,
		Name:     form.Name,
		Admin:    false,
		Verify:   true,
	}
	if err := mgm.Coll(data).Create(data); err != nil {
		return nil, err
	}
	return data, nil
}

// Signin defined signin user
func (h *Auth) Signin(form *models.Signin) (*models.Auth, error) {
	user := &models.User{}
	if err := mgm.Coll(user).First(bson.M{"email": form.Email}, user); err != nil {
		return nil, err
	}

	if user.Email == "" {
		return nil, fmt.Errorf("Email not found: %s", user.Email)
	}

	if err := common.PasswordCompare([]byte(form.Password), []byte(user.Password)); err != nil {
		return nil, err
	}

	token, err := createToken(user)
	if err != nil {
		return nil, err
	}

	return &models.Auth{
		UserID:      user.ID.Hex(),
		Email:       user.Email,
		Name:        user.Name,
		AccessToken: token,
	}, nil
}
