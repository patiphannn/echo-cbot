package common

import (
	"context"
	"os"
	"time"

	"github.com/Kamva/mgm/v2"
	"github.com/polnoy/echo-cbot/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

// ConnectDb is connect mongodb database
func ConnectDb() error {
	defer createIndex()

	mongo := os.Getenv("MONGO_HOST")
	if mongo == "" {
		mongo = "mongodb://localhost:27017"
	}

	db := os.Getenv("MONGO_DB_NAME")
	if db == "" {
		db = "echo-cbot"
	}

	//  _ = mgm.SetDefaultConfig(&mgm.Config{CtxTimeout: 12 * time.Second}, "go-book", options.Client().ApplyURI("mongodb://root:12345@localhost:27017"))
	err := mgm.SetDefaultConfig(&mgm.Config{CtxTimeout: 12 * time.Second}, db, options.Client().ApplyURI(mongo))
	if err != nil {
		panic(err.Error())
	}

	return err
}

func createIndex() error {
	// Create index
	coll := mgm.Coll(&models.User{})
	if _, err := coll.Indexes().CreateMany(
		context.Background(),
		[]mongo.IndexModel{
			{
				Keys:    bson.M{"email": 1},
				Options: options.Index().SetUnique(true),
			},
		},
	); err != nil {
		return err
	}
	return nil
}
