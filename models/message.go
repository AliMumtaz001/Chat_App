package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	// ID         int                `bson:"user_id" json:"user_id"`
	SenderID   int64              `bson:"sender_id" json:"sender_id"`
	ReceiverID string             `bson:"reciever_id" json:"reciever_id"`
	Content    string             `bson:"content" json:"content"`
	ChatID     primitive.ObjectID `bson:"chat_id" json:"chat_id"`
	Timestamp  time.Time          `bson:"time_stamp" json:"timestamp"`
}
