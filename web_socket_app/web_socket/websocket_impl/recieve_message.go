package websocket_impl

import (
	"log"

	"github.com/AliMumtazDev/socket/client"
	// upgradeconn "github.com/AliMumtazDev/Go_Chat_App/web_socket_app/database"
)

func (ws *WebSocketServiceImpl) ReceiveMessage(client *client.Client) ([]byte, error) {
	_, message, err := client.Conn.ReadMessage()
	if err != nil {
		log.Printf("Failed to read message from %s: %v", client.UserID, err)
		return nil, err
	}
	return message, nil
}
