package socket

import "github.com/gin-gonic/gin"

type WebSocketService interface {
	SendMessage(client *Client, message []byte) error
	ReceiveMessage(client *Client) ([]byte, error)
	RegisterWebSocketRoute(c *gin.Context)
}
