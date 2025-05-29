package userserviceimpl

import (
	"fmt"
	"log"

	"github.com/AliMumtazDev/Go_Chat_App/models"
	connection "github.com/AliMumtazDev/Go_Chat_App/socket_clint"
	"github.com/gin-gonic/gin"
)

func (s *UserServiceImpl) SendMessageService(c *gin.Context, message *models.Message) error {
	err := s.messageAuth.SendMessagedb(c, message)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}
	if connection.Conn == nil {
		return fmt.Errorf("WebSocket connection not established")
	}
	log.Print("recieving and sending ids are: ",message.ReceiverID, message.SenderID)
	messageToSend := models.ServerMesageToSocket{
		Action:        "send",
		DestinationID: int(message.ReceiverID),
		Content:       message.Content,
	}
	err = connection.Conn.WriteJSON(messageToSend)
	if err != nil {
		log.Println("Error writing to WebSocket server:", err)
		return err
	}
	log.Println("Message sent to WebSocket server:", messageToSend)

	return nil
}
