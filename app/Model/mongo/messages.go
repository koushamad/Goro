package mongo

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type (
	Messages struct {
		Id        bson.ObjectId `bson:"_id,omitempty" json:"_id"`
		ProfileId uint          `bson:"profileId"`
		Content   string        `bson:"content"`
		Type      int           `bson:"type"`
		Delivered int64         `bson:"delivered" json:"delivered"`
	}
)

func (message *Messages) GetId() bson.ObjectId {
	return message.Id
}

func (message *Messages) SetId(id bson.ObjectId) {
	message.Id = id
}

func (message *Messages) Add() error {
	message.Delivered = time.Now().Unix()
	return mongo.Collection("messages").Save(message)
}
