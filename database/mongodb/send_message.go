package mongodb

import (
	"context"
	"time"

	"github.com/AliMumtazDev/Go_Chat_App/models"
	"github.com/gin-gonic/gin"
)

func (r *StorageMongoImpl) SendMessagedb(c *gin.Context, msg *models.Message) error {

	message := models.Message{
		SenderID:   msg.SenderID,
		ReceiverID: msg.ReceiverID,
		Content:    msg.Content,
		Timestamp:  time.Now(),
	}
	collection := r.mongoClient.Database("chatdb").Collection("sendmsg")
	_, err := collection.InsertOne(context.TODO(), message)
	return err
}
