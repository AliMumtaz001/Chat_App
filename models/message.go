package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	SenderID   int64              `bson:"sender_id" json:"sender_id"`
	ReceiverID int64              `bson:"reciever_id,omitempty" json:"reciever_id,omitempty"`
	Content    string             `bson:"content" json:"content"`
	Timestamp  time.Time          `bson:"time_stamp" json:"timestamp"`
}
