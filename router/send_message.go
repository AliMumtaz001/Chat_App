package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AliMumtazDev/Go_Chat_App/models"
	socket "github.com/AliMumtazDev/Go_Chat_App/web_socket"
	"github.com/gin-gonic/gin"
)

// SendMessagereq godoc
// @Summary      Send a message
// @Description  Send a message from authenticated user to another user
// @Tags         messages
// @Accept       json
// @Produce      json
// @Param        message  body      models.Message  true  "Message data"
// @Success      200
// @Failure      400
// @Failure      401
// @Security     BearerAuth
// @Router       /sendmessage [post]
func (r *Router) SendMessagereq(c *gin.Context) {
	var message models.Message
	userID := c.MustGet("userID").(string)

	// Bind JSON request to the message model
	if err := c.ShouldBindJSON(&message); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Validate receiver ID
	if message.ReceiverID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receiver ID"})
		return
	}

	// Parse sender ID
	senderID, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		fmt.Println("Error parsing sender ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sender ID"})
		return
	}

	// Create a database message object
	dbMessage := models.Message{
		SenderID:   senderID,
		ReceiverID: message.ReceiverID,
		Content:    message.Content,
	}

	// Save the message to the database
	if err := r.UserService.SendMessageservice(c, senderID, message.ReceiverID, dbMessage); err != nil {
		fmt.Println("Error saving message:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Notify the recipient via WebSocket
	r.NotifyRecipient(userID, message)

	// Respond with success
	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully", "id": userID})
}

// notifyRecipient handles notifying the recipient via WebSocket
func (r *Router) NotifyRecipient(userID string, message models.Message) {
	socket.ConnMutex.Lock()
	defer socket.ConnMutex.Unlock()

	recipientID := strconv.FormatInt(message.ReceiverID, 10)
	fmt.Printf("Attempting to notify recipient: %s\n", recipientID)

	recipientConn, exists := socket.Connections[recipientID]
	if !exists {
		fmt.Printf("Recipient %s is not connected via WebSocket\n", recipientID)
		return
	}

	// Create a WebSocket message
	wsMessage := models.WebSocketMessage{
		Type:    "sendmessage",
		From:    userID,
		To:      recipientID,
		Content: message.Content,
	}

	// Marshal the WebSocket message
	msgBytes, err := json.Marshal(wsMessage)
	if err != nil {
		fmt.Printf("Failed to marshal WebSocket message: %v\n", err)
		return
	}

	// Send the WebSocket message
	if err := r.WebSocket.SendMessage(recipientConn, msgBytes); err != nil {
		fmt.Printf("Failed to send WebSocket message to recipient %s: %v\n", recipientID, err)
	} else {
		fmt.Printf("Message sent to recipient %s via WebSocket\n", recipientID)
	}
}
