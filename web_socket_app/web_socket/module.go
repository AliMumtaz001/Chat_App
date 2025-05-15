package socketinterface

import (
	socketimpl "github.com/AliMumtazDev/socket/websocket_impl"
	"github.com/gin-gonic/gin"
)

type WebSocketService interface {
	SendMessage(client *socketimpl.Client, message []byte) error
	ReceiveMessage(client *socketimpl.Client) ([]byte, error)
	RegisterWebSocketRoute(c *gin.Context)
}
