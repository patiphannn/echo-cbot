package models

import (
	"github.com/Kamva/mgm/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

// Intent defined intent object
type Intent struct {
	mgm.DefaultModel `bson:",inline"`

	Project     primitive.ObjectID `json:"project" bson:"project" binding:"required"`
	Name        string             `json:"name" bson:"name" binding:"required"`
	Context     string             `json:"context" bson:"context" binding:"required"`
	Question    []string           `json:"question" bson:"question" binding:"required"`
	Answer      []IAnswer          `json:"answer" bson:"answer" binding:"required"`
	Extra       bson.M             `json:"extra" bson:"extra"`
	Fallback    bool               `json:"fallback" bson:"fallback"`
	Fulfillment bool               `json:"fulfillment" bson:"fulfillment"`
	Welcome     bool               `json:"welcome" bson:"welcome"`
}

// IAnswer defined answer object
type IAnswer struct {
	Message string `json:"message" form:"message" bson:"message" binding:"required"`
	Extra   bson.M `json:"extra" form:"extra" bson:"extra"`
}

// RIntent defined intent request object
type RIntent struct {
	Name        string    `json:"name" form:"name" bson:"name" binding:"required"`
	Context     string    `json:"context" form:"context" bson:"context" binding:"required"`
	Question    []string  `json:"question" form:"question" bson:"question" binding:"required"`
	Answer      []IAnswer `json:"answer" form:"answer" bson:"answer" binding:"required"`
	Extra       bson.M    `json:"extra" form:"extra" bson:"extra"`
	Fallback    bool      `json:"fallback" form:"fallback" bson:"fallback"`
	Fulfillment bool      `json:"fulfillment" form:"fallback" bson:"fulfillment"`
	Welcome     bool      `json:"welcome" form:"welcome" bson:"welcome"`
}
