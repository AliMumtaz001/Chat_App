package socketimpl

import (
	"log"
	// upgradeconn "github.com/AliMumtazDev/Go_Chat_App/web_socket_app/database"
)

func (ws *WebSocketServiceImpl) ReceiveMessage(client *Client) ([]byte, error) {
	_, message, err := client.Conn.ReadMessage()
	if err != nil {
		log.Printf("Failed to read message from %s: %v", client.UserID, err)
		return nil, err
	}
	return message, nil
}
