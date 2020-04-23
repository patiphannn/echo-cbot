package models

import (
	"gopkg.in/mgo.v2/bson"
)

// Response defined reponse object
type Response struct {
	Question string `json:"question" form:"question" bson:"question" binding:"required"`
	Options  bson.M `json:"options" form:"options" bson:"options"`
	Extra    bson.M `json:"extra" form:"extra" bson:"extra"`
}

// Answer defined answer object
type Answer struct {
	Message string `json:"message" bson:"message" binding:"required"`
}

// Bot defined bot object
type Bot struct {
	Name        string   `json:"name" bson:"name"`
	Fulfillment string   `json:"fulfillment" bson:"fulfillment"`
	Fallback    int      `json:"fallback" bson:"fallback"`
	Intents     []Intent `json:"intents" bson:"intents"`
}
