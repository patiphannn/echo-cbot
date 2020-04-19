package models

import (
	"github.com/Kamva/mgm/v2"
)

// Project defined project object
type Project struct {
	mgm.DefaultModel `bson:",inline"`

	Name        string `json:"name" form:"name" bson:"name" binding:"required"`
	Fulfillment string `json:"fulfillment" form:"fulfillment" bson:"fulfillment" binding:"required"`
	Enabled     bool   `json:"enabled" form:"enabled" bson:"enabled" binding:"required"`
	Key         string `json:"key" bson:"key"`
	User        string `json:"user" bson:"user"`
}

// RProject defined project request object
type RProject struct {
	Name        string `json:"name" form:"name" bson:"name" binding:"required"`
	Fulfillment string `json:"fulfillment" form:"fulfillment" bson:"fulfillment" binding:"required"`
	Enabled     bool   `json:"enabled" form:"enabled" bson:"enabled" binding:"required"`
}
