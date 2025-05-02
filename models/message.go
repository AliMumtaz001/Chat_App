package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	// ID         primitive.ObjectID `bson:"_id" json:"id"`
	SenderID   string             `bson:"sender_id" json:"sender_id"`
	ReceiverID string             `bson:"reciever_id" json:"reciever_id"`
	Content    string             `bson:"content" json:"content"`
	ChatID     primitive.ObjectID `bson:"chat_id" json:"chat_id"`
	Timestamp  time.Time          `bson:"time_stamp" json:"timestamp"`
}
