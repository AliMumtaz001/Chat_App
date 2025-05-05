package userserviceimpl

import (
	"fmt"

	"github.com/AliMumtaz001/Go_Chat_App/models"
	"github.com/gin-gonic/gin"
)

func (s *UserServiceImpl) SendMessageservice(c *gin.Context, senderID, receiverID string, message models.Message) error {
	// recID, err := strconv.ParseInt(receiverID, 10, 64)
	err := s.messageAuth.SendMessagedb(c, senderID, receiverID, message)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}
	return nil
}

func (s *UserServiceImpl) GetMessageservice(c *gin.Context, senderID, receiverID string) ([]models.Message, error) {
	messages, err := s.messageAuth.GetMessagedb(c, senderID, receiverID)
	if err != nil {
		return nil, fmt.Errorf("failed to get messages: %w", err)
	}
	return messages, nil
}
