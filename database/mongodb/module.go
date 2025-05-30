package mongodb

import (
	"github.com/AliMumtazDev/Go_Chat_App/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Storage interface {
	SaveMessage(c *gin.Context, msg *models.Message) error
	GetMessagedb(c *gin.Context, senderID string, receiverID string) ([]models.Message, error)
	UpdateMessagedb(c *gin.Context, messageID primitive.ObjectID, message models.Message) error
	DeleteMessagedb(c *gin.Context, messageID primitive.ObjectID) error
}
