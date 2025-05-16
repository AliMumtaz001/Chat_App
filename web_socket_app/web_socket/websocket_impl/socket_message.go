package websocket_impl

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (impl *WebSocketServiceImpl) RegisterWebSocketRoute(c *gin.Context) {
	// ws := websocket_impl.NewWebSocketService()
	//NewWebSocketImpl()
	client, err := impl.WSImpl.UpgradeConnection(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upgrade connection"})
		return
	}
	for {
		message, err := impl.ReceiveMessage(client)
		if err != nil {
			log.Printf("Error reading message: %v", err)
			return
		}
		err = impl.SendMessage(client, message)
		if err != nil {
			log.Printf("Error sending message: %v", err)
			return
		}
	}
}
