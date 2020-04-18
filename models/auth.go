package models

import (
	"github.com/Kamva/mgm/v2"
)

// User defined user model
type User struct {
	mgm.DefaultModel `bson:",inline"`

	Email    string `json:"email" bson:"email" validate:"required,email"`
	Password string `json:"-" bson:"password" validate:"required,gte=6,lte=20"`
	Name     string `json:"name" bson:"name" validate:"required"`
	Admin    bool   `json:"admin" bson:"admin" validate:"required"`
	Verify   bool   `json:"verify" bson:"verify" validate:"required"`
}

// Auth defined auth object
type Auth struct {
	UserID      string `json:"_id" bson:"_id" validate:"required"`
	Email       string `json:"email" bson:"email" validate:"required,email"`
	Name        string `json:"name" bson:"name" validate:"required"`
	AccessToken string `json:"access_token" bson:"access_token" validate:"required"`
}

// Signup defined signup object
type Signup struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,gte=6,lte=20"`
	Confirm  string `json:"confirm" form:"confirm" validate:"eqfield=Password"`
	Name     string `json:"name" form:"name" validate:"required"`
}

// Signin defined signin object
type Signin struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,gte=6,lte=20"`
}
