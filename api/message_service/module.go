package userserviceimpl

import (
	"github.com/AliMumtaz001/Go_Chat_App/models"
	"github.com/gin-gonic/gin"
)

type UserService interface {
	SendMessageservice(c *gin.Context, senderID, receiverID int64, message models.Message) error
	GetMessageservice(c *gin.Context, senderID string, receiverID string) ([]models.Message, error)
	UpdateMessageservice(c *gin.Context, messageID string, message models.Message, userID string) error
	DeleteMessageservice(c *gin.Context, messageID string, userID string) error
}
