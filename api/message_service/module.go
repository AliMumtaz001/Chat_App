package userserviceimpl

import (
	"github.com/AliMumtazDev/Go_Chat_App/models"
	"github.com/gin-gonic/gin"
)

type UserService interface {
	SendMessageService(c *gin.Context, message *models.Message) error
	GetMessageService(c *gin.Context, senderID string, receiverID string) ([]models.Message, error)
	UpdateMessageService(c *gin.Context, messageID string, message models.Message) error
	DeleteMessageService(c *gin.Context, messageID string) error
}
