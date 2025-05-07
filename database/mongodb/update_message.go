package mongodb

import (
	"fmt"

	"github.com/AliMumtaz001/Go_Chat_App/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *StorageMongoImpl) UpdateMessagedb(c *gin.Context, messageID primitive.ObjectID, message models.Message) error {
	if s.mongoClient == nil {
		return fmt.Errorf("mongo client is not initialized")
	}

	db := s.mongoClient.Database("chatdb")
	collection := db.Collection("sendmsg")

	filter := bson.M{"_id": messageID}

	update := bson.M{
		"$set": bson.M{
			"content": message.Content,
		},
	}

	result, err := collection.UpdateOne(c, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update message with ID %v: %w", messageID, err)
	}

	if result == nil {
		return fmt.Errorf("no message found with ID %v", messageID)
	}

	return nil
}
