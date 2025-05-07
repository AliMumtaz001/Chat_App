package mongodb

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func (m *StorageMongoImpl) DeleteMessagedb(c *gin.Context, messageID primitive.ObjectID) error {
	if m.mongoClient == nil {
		return fmt.Errorf("mongo client is not initialized")
	}

	db := m.mongoClient.Database("chatdb")
	collection := db.Collection("sendmsg")

	filter := bson.M{
		"_id": messageID,
	}

	result, err := collection.DeleteOne(c, filter)
	if err != nil {
		return err
	}
	fmt.Println("Result:", result)

	if result == nil {
		return fmt.Errorf("no message found with the given IDsss")
	}

	return nil
}
