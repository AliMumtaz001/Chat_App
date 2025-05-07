package mongodb

import (
	"context"
	"time"

	"github.com/AliMumtaz001/Go_Chat_App/models"
	"github.com/gin-gonic/gin"
)

func (r *StorageMongoImpl) SendMessagedb(c *gin.Context, senderID, receiverID int64, msg models.Message) error {

	message := models.Message{
		SenderID:   senderID,
		ReceiverID: receiverID,
		Content:    msg.Content,
		Timestamp:  time.Now(),
	}
	collection := r.mongoClient.Database("chatdb").Collection("sendmsg")
	_, err := collection.InsertOne(context.TODO(), message)
	return err
}
