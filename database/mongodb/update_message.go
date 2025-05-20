package mongodb

import (
	"fmt"
	"strconv"

	"github.com/AliMumtazDev/Go_Chat_App/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *StorageMongoImpl) UpdateMessagedb(c *gin.Context, messageID primitive.ObjectID, message models.Message) error {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(401, gin.H{"error": "User not authenticated"})
		return fmt.Errorf("user not authenticated")
	}
	if s.mongoClient == nil {
		return fmt.Errorf("mongo client is not initialized")
	}
	
	collection := s.mongoClient.Database("chatdb").Collection("sendmsg")
	var doc bson.M
	err := collection.FindOne(c, bson.M{"_id": messageID}).Decode(&doc)
	if err == mongo.ErrNoDocuments {
		return fmt.Errorf("document not found for _id: %v", messageID)
	}
	if err != nil {
		return fmt.Errorf("error checking document: %v", err)
	}
	userIDStr, ok := userID.(string)
	if !ok {
		return fmt.Errorf("userID is not a string")
	}
	userIDInt, err := strconv.Atoi(userIDStr)
	if err != nil {
		userIDInt = 0
	}
	filter := bson.M{
		"_id": messageID,
		"$or": []bson.M{
			{
				"sender_id": userIDInt,
			},
			{
				"receiver_id": userIDInt,
			},
		},
	}
	update := bson.M{
		"$set": bson.M{
			"content": message.Content,
		},
	}
	result, err := collection.UpdateOne(c, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update message with ID %v: %w", messageID, err)
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("no message found with the given ID or user is not authorized")
	}
	return nil
}
