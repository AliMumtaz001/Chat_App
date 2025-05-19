package routes

import (
	"log"
	"net/http"

	"github.com/AliMumtazDev/socket/client"
	"github.com/AliMumtazDev/socket/models"
	"github.com/gin-gonic/gin"
)

func (r *SocketRouter) RegisterWebSocketRoute(c *gin.Context) {
	log.Println("Inside RegisterWebSocketRoute")
	userID := c.MustGet("userID").(string)
	conn, err := models.Upgrader.Upgrade(c.Writer, c.Request, nil)
	client := &client.Client{Conn: conn, UserID: userID}
	err = r.WebSocket.AddConn(userID, client, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upgrade connection"})
		return
	}
}
