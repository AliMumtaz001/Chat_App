package socketinterface

import (
	"github.com/AliMumtazDev/socket/client"
	"github.com/gin-gonic/gin"
)

type WebSocketService interface {
	SendMessage(client *client.Client, message []byte) error
	ReceiveMessage(client *client.Client) ([]byte, error)
	RegisterWebSocketRoute(c *gin.Context)
}
