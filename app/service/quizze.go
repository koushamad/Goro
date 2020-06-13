package service

import (
	"github.com/koushamad/goro/app/Model/mongo"
	"strconv"
)

type (
	Quiz struct {
		QuizId string `json:"quizId"`
	}

	MessageSend struct {
		QuizId  string `json:"quizId"`
		Content string `json:"content"`
		Type    string `json:"type"`
	}

	MessageReceive struct {
		QuizId    string `json:"quizId"`
		Content   string `json:"content"`
		Type      int    `json:"type"`
		ProfileId uint   `json:"profileId"`
		Delivered int64  `json:"delivered"`
	}
)

func (quiz Quiz) Join() bool {
	quizzes := mongo.Quizzes{}
	quizzes.GetById(quiz.QuizId)
	return quizzes.Id.String() == quiz.QuizId
}

func (quiz Quiz) Leave() bool {
	quizzes := mongo.Quizzes{}
	quizzes.GetById(quiz.QuizId)
	return quizzes.Id.String() == quiz.QuizId
}

func (mes MessageSend) Send(profileId uint) (error, MessageReceive) {
	quiz := mongo.Quizzes{}
	quiz.GetById(mes.QuizId)
	t, _ := strconv.Atoi(mes.Type)
	message := mongo.Messages{
		ProfileId: profileId,
		Content:   mes.Content,
		Type:      t,
	}
	return quiz.AddMessage(&message), MessageReceive{
		QuizId:    mes.QuizId,
		Content:   message.Content,
		Type:      message.Type,
		ProfileId: message.ProfileId,
		Delivered: message.Delivered,
	}
}
