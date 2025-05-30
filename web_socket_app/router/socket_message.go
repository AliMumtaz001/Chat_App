package routes

import (
	"log"
	"net/http"

	"github.com/AliMumtazDev/socket/models"
	"github.com/gin-gonic/gin"
)

func (r *SocketRouter) RegisterWebSocketRoute(c *gin.Context) {
	log.Println("Inside RegisterWebSocketRoute")
	userID := c.MustGet("userID").(string)
	conn, err := models.Upgrader.Upgrade(c.Writer, c.Request, nil)
	go func() {
		err = r.WebSocket.AddConn(userID, conn, c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upgrade connection"})
			return
		}

	}()
}

func (r *SocketRouter) SaveBackendConnection(c *gin.Context) {
	userID := "-1"
	wsConn, err := models.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("WebSocket Upgrade failed: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upgrade connection"})
		return
	}
	log.Printf("Backend WebSocket connection established for userID: %s", userID)
	go func() {
		err = r.WebSocket.AddConn(userID, wsConn, c)
		if err != nil {
			log.Printf("AddConn error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register connection"})
			return
		}
	}()

}
