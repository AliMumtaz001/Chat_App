package userserviceimpl

import (
	"fmt"

	"github.com/AliMumtazDev/Go_Chat_App/models"
	"github.com/gin-gonic/gin"
)

func (s *UserServiceImpl) SendMessageservice(c *gin.Context, senderID, receiverID int64, message models.Message) error {
	err := s.messageAuth.SendMessagedb(c, senderID, receiverID, message)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}
	return nil
}
