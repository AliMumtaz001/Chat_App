package websocket_impl

import (
	"net/http"

	"github.com/AliMumtazDev/socket/client"
	"github.com/AliMumtazDev/socket/models"
	"github.com/gin-gonic/gin"
)

func (impl *WebSocketServiceImpl) RegisterWebSocketRoute(c *gin.Context) {
	userID := c.MustGet("userID").(string)
	conn, err := models.Upgrader.Upgrade(c.Writer, c.Request, nil)
	client := &client.Client{Conn: conn, UserID: userID}
	err = impl.AddConn(userID, client, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upgrade connection"})
		return
	}
}
