package mongo

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	Quizzes struct {
		Id       bson.ObjectId `bson:"_id,omitempty" json:"_id"`
		Messages []string      `bson:"messages"`
	}
)

func (quiz *Quizzes) GetId() bson.ObjectId {
	return quiz.Id
}

func (quiz *Quizzes) SetId(id bson.ObjectId) {
	quiz.Id = id
}

func (quiz *Quizzes) GetById(QuizId string) {
	mongo.Collection("quizzes").FindById(bson.ObjectIdHex(QuizId), quiz)
}

func (quiz *Quizzes) AddMessage(message *Messages) error {
	err := message.Add()
	if err == nil {
		quiz.Messages = append(quiz.Messages, message.Id.Hex())
		err = quiz.Save()
	}
	return err
}

func (quiz *Quizzes) Save() error {
	return mongo.Collection("quizzes").Save(quiz)
}
