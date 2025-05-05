package mongodb

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/AliMumtaz001/Go_Chat_App/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StorageMongoImpl struct {
	mongoClient *mongo.Client
}

func NewStorage(input *mongo.Client) Storage {
	return &StorageMongoImpl{
		mongoClient: input,
	}
}

func (r *StorageMongoImpl) SendMessagedb(c *gin.Context, senderID, receiverID string, msg models.Message) error {
	senID, err := strconv.ParseInt(senderID, 10, 64)
	recID, err := strconv.ParseInt(receiverID, 10, 64)

	message := models.Message{
		SenderID:   senID,
		ReceiverID: recID,
		Content:    msg.Content,
		Timestamp:  time.Now(),
	}
	collection := r.mongoClient.Database("chatdb").Collection("sendmsg")
	_, err = collection.InsertOne(context.TODO(), message)
	return err
}

func (r *StorageMongoImpl) GetMessagedb(c *gin.Context, senderID string, receiverID string) ([]models.Message, error) {
    filter := bson.M{}

    // Convert senderID and receiverID to integers
    senID, err := strconv.ParseInt(senderID, 10, 64)
    if err == nil {
        filter["sender_id"] = senID
    }
    recID, err := strconv.ParseInt(receiverID, 10, 64)
    if err == nil {
        filter["reciever_id"] = recID // Corrected spelling
    }

    // Debug log for the filter
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
        fmt.Printf("Decoded message: %+v\n", message) // Debug log
        messages = append(messages, message)
    }

    return messages, nil
}