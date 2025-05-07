package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) GetMessagereq(c *gin.Context) {

	senderID := c.Query("sender_id")
	receiverID := c.Query("reciever_id")

	if senderID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing sender_id "})
		return
	}
	if receiverID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing reciever_id"})
		return
	}

	fmt.Printf("Sender ID: %s, Receiver ID: %s\n", senderID, receiverID)

	messages, err := r.UserService.GetMessageservice(c, senderID, receiverID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, messages)
}
