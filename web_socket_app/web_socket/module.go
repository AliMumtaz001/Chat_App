package web_socket

import (
	"github.com/AliMumtazDev/socket/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebSocketService interface {
	SendMessage(client *models.Client, message []byte) error
	ReceiveMessage(client *models.Client) ([]byte, error)
	AddConn(userID string, client *websocket.Conn, c *gin.Context) error
}
