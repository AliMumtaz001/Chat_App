package userserviceimpl

import (
	"fmt"

	"github.com/AliMumtazDev/Go_Chat_App/models"
	"github.com/gin-gonic/gin"
)

func (s *UserServiceImpl) GetMessageservice(c *gin.Context, senderID, receiverID string) ([]models.Message, error) {
	messages, err := s.messageAuth.GetMessagedb(c, senderID, receiverID)
	if err != nil {
		return nil, fmt.Errorf("failed to get messages: %w", err)
	}
	return messages, nil
}
