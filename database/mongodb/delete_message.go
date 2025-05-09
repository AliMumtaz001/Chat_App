package mongodb
import (
	"fmt"
	"strconv"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)
func (m *StorageMongoImpl) DeleteMessagedb(c *gin.Context, messageID primitive.ObjectID, userID string) error {
	if m.mongoClient == nil {
		return fmt.Errorf("mongo client is not initialized")
	}
	db := m.mongoClient.Database("chatdb")
	collection := db.Collection("sendmsg")
	var doc bson.M
	err := collection.FindOne(c, bson.M{"_id": messageID}).Decode(&doc)
	if err == mongo.ErrNoDocuments {
		return fmt.Errorf("Document not found for _id:", messageID)
	}
	if err != nil {
		return fmt.Errorf("error checking document: %v", err)
	}
	userIDInt, err := strconv.Atoi(userID)
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
	result, err := collection.DeleteOne(c, filter)
	if err != nil {
		return fmt.Errorf("delete error: %v", err)
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("no message found with the given ID or user is not authorized")
	}
	return nil
	
}