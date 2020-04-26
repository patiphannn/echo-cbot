package common

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Kamva/mgm/v2"
	"github.com/polnoy/echo-cbot/models"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func mongoHost() string {
	mongo := os.Getenv("MONGO_HOST")
	if mongo != "" {
		return mongo
	}

	mongoV := viper.GetString("mongo_host")
	if mongoV != "" {
		return mongoV
	}

	return "mongodb://localhost:27017"
}

func mongoDB() string {
	mongo := os.Getenv("MONGO_DB_NAME")
	if mongo != "" {
		return mongo
	}

	mongoV := viper.GetString("mongo_db")
	if mongoV != "" {
		return mongoV
	}

	return "mongodb://localhost:27017"
}

// ConnectDb is connect mongodb database
func ConnectDb() error {
	defer createIndex()
	defer seed()

	mongo := mongoHost()

	db := mongoDB()

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

func seed() error {
	seed := viper.GetString("seed")
	if seed != "true" {
		return nil
	}

	fmt.Println("Seed: ")

	// Users
	user, err := userSeed()
	if err != nil {
		return err
	}
	userID := user.ID
	fmt.Println("userID: ", userID)

	// Projects
	project, err := projectSeed(userID)
	if err != nil {
		return err
	}
	projID := project.ID
	fmt.Println("projID: ", projID)

	// Intents
	if err := intentSeed(projID); err != nil {
		return err
	}

	return nil
}

func userSeed() (*models.User, error) {
	// Drop User Collection
	mgm.CollectionByName("users").Drop(mgm.Ctx())

	// Create User
	password, err := GeneratePasswordHash([]byte("adminbot"))
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:    "admin@admin.com",
		Password: password,
		Name:     "Admin",
		Admin:    true,
		Verify:   true,
	}
	if err := mgm.Coll(user).Create(user); err != nil {
		fmt.Println("err: ", err.Error())
		return nil, err
	}

	return user, nil
}

func projectSeed(userID primitive.ObjectID) (*models.Project, error) {
	// Drop Project Collection
	mgm.CollectionByName("projects").Drop(mgm.Ctx())

	// Create Project
	project := &models.Project{
		Name:        "cBot",
		Fulfillment: "",
		Enabled:     true,
		User:        userID,
	}
	coll := mgm.Coll(project)
	if err := coll.Create(project); err != nil {
		fmt.Println("err: ", err.Error())
		return nil, err
	}
	key, err := CreateKey(project)
	if err != nil {
		return nil, err
	}
	project.Key = key
	if err := coll.Update(project); err != nil {
		return nil, err
	}

	return project, nil
}

func intentSeed(projID primitive.ObjectID) error {
	// Drop Intent Collection
	mgm.CollectionByName("intents").Drop(mgm.Ctx())

	// Create Intent
	// Fallback
	{
		question := []string{}
		answer := []models.IAnswer{
			{
				Message: "บอทงงครับ",
				Extra:   bson.M{},
			},
			{
				Message: "งงครับผม",
				Extra:   bson.M{},
			},
			{
				Message: "อะไรนะครับ",
				Extra:   bson.M{},
			},
		}
		data := &models.Intent{
			Project:     projID,
			Name:        "Fallback intent",
			Context:     "",
			Question:    question,
			Answer:      answer,
			Extra:       bson.M{},
			Fallback:    true,
			Fulfillment: false,
			Welcome:     false,
		}
		if err := mgm.Coll(data).Create(data); err != nil {
			fmt.Println("err: ", err.Error())
			return err
		}
	}

	// ทักทาย
	{
		question := []string{"สวัสดี", "ดีครับ", "Hello", "Hi"}
		answer := []models.IAnswer{
			{
				Message: "สวัสดีครับ",
				Extra:   bson.M{},
			},
			{
				Message: "ดีครับ",
				Extra:   bson.M{},
			},
			{
				Message: "Bello",
				Extra:   bson.M{},
			},
		}
		data := &models.Intent{
			Project:     projID,
			Name:        "ทักทาย",
			Context:     "",
			Question:    question,
			Answer:      answer,
			Extra:       bson.M{},
			Fallback:    false,
			Fulfillment: false,
			Welcome:     false,
		}
		if err := mgm.Coll(data).Create(data); err != nil {
			fmt.Println("err: ", err.Error())
			return err
		}
	}

	// ดี
	{
		question := []string{"ดีจัง", "ดีสุดๆ", "ยอดไปเลย", "ดีมาก", "ดีๆ"}
		answer := []models.IAnswer{
			{
				Message: "ยินดีที่ได้ช่วยเหลือครับ",
				Extra:   bson.M{},
			},
			{
				Message: "เจ๋งใช่มะ",
				Extra:   bson.M{},
			},
			{
				Message: "ดีใช่มะ",
				Extra:   bson.M{},
			},
		}
		data := &models.Intent{
			Project:     projID,
			Name:        "ดี",
			Context:     "",
			Question:    question,
			Answer:      answer,
			Extra:       bson.M{},
			Fallback:    false,
			Fulfillment: false,
			Welcome:     false,
		}
		if err := mgm.Coll(data).Create(data); err != nil {
			fmt.Println("err: ", err.Error())
			return err
		}
	}

	// ทานข้าว
	{
		question := []string{"กินข้าวยัง", "กินข้าวหรือยัง"}
		answer := []models.IAnswer{
			{
				Message: "ยังไม่หิวเลยครับ",
				Extra:   bson.M{},
			},
			{
				Message: "ผมเป็นบอทไม่ทานข้าวครับ",
				Extra:   bson.M{},
			},
			{
				Message: "ทานก่อนได้เลยครับ",
				Extra:   bson.M{},
			},
		}
		data := &models.Intent{
			Project:     projID,
			Name:        "กินข้าวหรือยัง",
			Context:     "",
			Question:    question,
			Answer:      answer,
			Extra:       bson.M{},
			Fallback:    false,
			Fulfillment: false,
			Welcome:     false,
		}
		if err := mgm.Coll(data).Create(data); err != nil {
			fmt.Println("err: ", err.Error())
			return err
		}
	}

	// โง่
	{
		question := []string{"บอทโง่", "โง่", "ง่าว", "ควาย", "ฟาย"}
		answer := []models.IAnswer{
			{
				Message: "ฮาๆ บอทยังละอ่อนครับ",
				Extra:   bson.M{},
			},
			{
				Message: "โง่แต่รักนะ",
				Extra:   bson.M{},
			},
			{
				Message: "คนสร้างบอทก็โง่ครับ",
				Extra:   bson.M{},
			},
		}
		data := &models.Intent{
			Project:     projID,
			Name:        "โง่",
			Context:     "",
			Question:    question,
			Answer:      answer,
			Extra:       bson.M{},
			Fallback:    false,
			Fulfillment: false,
			Welcome:     false,
		}
		if err := mgm.Coll(data).Create(data); err != nil {
			fmt.Println("err: ", err.Error())
			return err
		}
	}

	return nil
}
