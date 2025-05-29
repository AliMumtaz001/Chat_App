package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetMessagereq godoc
// @Summary      Get messages
// @Description  Retrieve messages between two users
// @Tags         messages
// @Produce      json
// @Param        sender_id    query  string  true  "Sender ID"
// @Param        reciever_id  query  string  true  "Receiver ID"
// @Success      200          {array}   models.Message
// @Failure      400
// @Failure      500
// @Security     BearerAuth
// @Router       /getmessage [get]
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

	messages, err := r.UserService.GetMessageService(c, senderID, receiverID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, messages)
}
