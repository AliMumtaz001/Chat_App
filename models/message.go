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

type WebSocketMessage struct {
	Type    string `json:"type"`    // "sendmessage" or "getmessage"
	Content string `json:"content"` // Message content
	To      string `json:"to"`      // Recipient ID
	From    string `json:"from"`    // Sender ID
}
