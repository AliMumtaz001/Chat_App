package websocket_impl

import (
	"log"

	"github.com/AliMumtazDev/socket/models"
)

func (ws *WebSocketServiceImpl) ReceiveMessage(client *models.Client) ([]byte, error) {
	_, message, err := client.Conn.ReadMessage()
	if err != nil {
		log.Printf("Failed to read message from %s: %v", client.UserID, err)
		return nil, err
	}
	return message, nil
}
