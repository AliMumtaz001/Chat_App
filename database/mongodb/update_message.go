package mongodb
import (
	"fmt"
	"strconv"
	"github.com/AliMumtaz001/Go_Chat_App/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)
func (s *StorageMongoImpl) UpdateMessagedb(c *gin.Context, messageID primitive.ObjectID, message models.Message, userID string) error {
	if s.mongoClient == nil {
		return fmt.Errorf("mongo client is not initialized")
	}
	db := s.mongoClient.Database("chatdb")
	collection := db.Collection("sendmsg")
	var doc bson.M
	err := collection.FindOne(c, bson.M{"_id": messageID}).Decode(&doc)
	if err == mongo.ErrNoDocuments {
		fmt.Println("Document not found for _id:", messageID)
		return fmt.Errorf("no message found with the given ID")
	}
	if err != nil {
		return fmt.Errorf("error checking document: %v", err)
	}
	fmt.Println("Document found:", doc)
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		userIDInt = 0 
	}
	filter := bson.M{
		"_id": messageID,
		"$or": []bson.M{
			{
				"sender_id": bson.M{
					"$in": []interface{}{userID, userIDInt},
				},
			},
			{
				"receiver_id": bson.M{
					"$in": []interface{}{userID, userIDInt},
				},
			},
		},
	}
	update := bson.M{
		"$set": bson.M{
			"content": message.Content,
		},
	}
	fmt.Println("Filter:", filter)
	result, err := collection.UpdateOne(c, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update message with ID %v: %w", messageID, err)
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("no message found with the given ID or user is not authorized")
	}
	fmt.Println("Result: MatchedCount:", result.MatchedCount, "ModifiedCount:", result.ModifiedCount)
	return nil
}