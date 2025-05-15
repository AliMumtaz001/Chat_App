package socket

import "log"

func (ws *WebSocketImpl) ReceiveMessage(client *Client) ([]byte, error) {
	_, message, err := client.Conn.ReadMessage()
	if err != nil {
		log.Printf("Failed to read message from %s: %v", client.UserID, err)
		return nil, err
	}
	return message, nil
}
