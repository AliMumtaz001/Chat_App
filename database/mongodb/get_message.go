package mongodb

import (
	"context"
	"fmt"
	"strconv"

	"github.com/AliMumtazDev/Go_Chat_App/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *StorageMongoImpl) GetMessagedb(c *gin.Context, senderID string, receiverID string) ([]models.Message, error) {
	filter := bson.M{}

	senID, err := strconv.ParseInt(senderID, 10, 64)
	if err == nil {
		filter["sender_id"] = senID
	}
	recID, err := strconv.ParseInt(receiverID, 10, 64)
	if err == nil {
		filter["reciever_id"] = recID
	}

	fmt.Printf("Constructed Filter: %+v\n", filter)

	collection := r.mongoClient.Database("chatdb").Collection("sendmsg")
	findOpt := options.Find()
	cursor, err := collection.Find(context.TODO(), filter, findOpt)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var messages []models.Message
	for cursor.Next(context.TODO()) {
		var message models.Message
		if err := cursor.Decode(&message); err != nil {
			return nil, err
		}
		fmt.Printf("Decoded message: %+v\n", message)
		messages = append(messages, message)
	}

	return messages, nil
}
