package services

import (
	"fmt"
	"math/rand"

	"github.com/Kamva/mgm/builder"
	"github.com/Kamva/mgm/v2"
	"github.com/Kamva/mgm/v2/operator"
	"github.com/cdipaolo/goml/base"
	naivebaye "github.com/cdipaolo/goml/text"
	"github.com/dgrijalva/jwt-go"
	gothaiwordcut "github.com/narongdejsrn/go-thaiwordcut"
	"github.com/polnoy/echo-cbot/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

// Response is all response services
type Response struct{}

func (h *Response) bot(id string, form *models.Response) (*models.Bot, error) {
	_id, _ := primitive.ObjectIDFromHex(id)

	result := []models.Bot{}
	intentColl := mgm.Coll(&models.Intent{}).Name()
	err := mgm.Coll(&models.Project{}).SimpleAggregate(
		&result,
		bson.M{operator.Match: bson.M{"_id": _id, "enabled": true}},
		builder.Lookup(intentColl, "_id", "project", "intents"),
		bson.M{operator.Unwind: bson.M{"path": "$intents"}},
		bson.M{operator.Group: bson.M{
			"_id":         "$_id",
			"name":        bson.M{"$first": "$name"},
			"fulfillment": bson.M{"$first": "$fulfillment"},
			"intents":     bson.M{"$push": "$intents"},
		}},
		bson.M{operator.Project: bson.M{
			"name":        "$name",
			"fulfillment": "$fulfillment",
			"fallback": bson.M{
				"$reduce": bson.M{
					"input":        "$intents",
					"initialValue": bson.M{"key": 0, "val": -1},
					"in": bson.M{
						"$cond": bson.M{
							"if": bson.M{
								"$eq": []interface{}{"$$this.fallback", true},
							},
							"then": bson.M{
								"key": bson.M{
									"$add": []interface{}{"$$value.key", 1},
								},
								"val": "$$value.key",
							},
							"else": bson.M{
								"key": bson.M{"$add": []interface{}{"$$value.key", 1}},
								"val": "$$value.val",
							},
						},
					},
				},
			},
			"intents": "$intents",
		}},
		bson.M{operator.Project: bson.M{
			"name":        "$name",
			"fulfillment": "$fulfillment",
			"fallback":    "$fallback.val",
			"intents":     "$intents",
		}},
	)
	if err != nil {
		return nil, err
	}

	return &result[0], nil
}

func (h *Response) naiveBaye(text string, bot *models.Bot) (uint8, float64, error) {
	segmenter := gothaiwordcut.Wordcut()
	segmenter.LoadDefaultDict()
	// create the channel of data and errors
	stream := make(chan base.TextDatapoint, 100)
	errors := make(chan error)
	len := uint8(len(bot.Intents))

	model := naivebaye.NewNaiveBayes(stream, len, base.OnlyWordsAndNumbers)

	go model.OnlineLearn(errors)

	for kintent, intent := range bot.Intents {
		questions := intent.Question
		fmt.Println("questions: ", kintent, questions)

		for _, question := range questions {
			wcut := segmenter.Segment(question)
			str := fmt.Sprintf("%v", wcut)
			fmt.Println("question: ", kintent, str)

			stream <- base.TextDatapoint{
				X: str,
				Y: uint8(kintent),
			}
		}
	}

	close(stream)

	for {
		err, _ := <-errors
		if err != nil {
			fmt.Printf("Error passed: %v", err)
		} else {
			// training is done!
			break
		}
	}

	// now you can predict like normal
	wcut := segmenter.Segment(text)
	str := fmt.Sprintf("%v", wcut)
	fmt.Println("text form: ", str)
	index, score := model.Probability(str) // 0
	return index, score, nil
}

func (h *Response) randomAnswer(index int, bot *models.Bot) string {
	intent := bot.Intents[index]
	answers := intent.Answer
	answer := answers[rand.Intn(len(answers))]

	return answer.Message
}

// Answer defined response answer.
func (h *Response) Answer(form *models.Response, profile jwt.MapClaims) (*models.Answer, error) {
	projectID := profile["_id"].(string)
	// Get intents
	bot, err := h.bot(projectID, form)
	if err != nil {
		return nil, err
	}
	// Naivebaye
	index, score, err := h.naiveBaye(form.Question, bot)
	if err != nil {
		return nil, err
	}
	fmt.Println("index: ", index)
	fmt.Println("score: ", score)

	var answer = ""
	if score > float64(0.5) {
		// match
		answer = h.randomAnswer(int(index), bot)
	} else {
		// fallback
		answer = h.randomAnswer(bot.Fallback, bot)
	}

	return &models.Answer{
		Message: answer,
	}, nil
}
