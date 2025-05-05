package mongodb

import (
	"github.com/AliMumtaz001/Go_Chat_App/models"
	"github.com/gin-gonic/gin"
)

type Storage interface {
	SendMessagedb(c *gin.Context, senderID, receiverID string, msg models.Message) error
	GetMessagedb(c *gin.Context, senderID string, receiverID string) ([]models.Message, error)
}
