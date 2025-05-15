package socketimpl

import (
	"log"

	"github.com/gorilla/websocket"
)

func (ws *WebSocketServiceImpl) SendMessage(client *Client, message []byte) error {
	err := client.Conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Printf("Failed to send message to %s: %v", client.UserID, err)
		return err
	}
	return nil
}
