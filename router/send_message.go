package routes
import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AliMumtazDev/Go_Chat_App/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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
	if err := c.ShouldBindJSON(&message); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	sID, _ := strconv.ParseInt(userID, 10, 64)
	if message.ReceiverID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receiver ID"})
		return
	}

	dbMessage := models.Message{
		SenderID:   sID,
		ReceiverID: message.ReceiverID,
		Content:    message.Content,
	}
	err := r.UserService.SendMessageservice(c, sID, message.ReceiverID, dbMessage)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	mutex.Lock()
	recipientID := strconv.FormatInt(message.ReceiverID, 10)
	fmt.Printf("Attempting to notify recipient: %s\n", recipientID)
	recipientConn, exists := connections[recipientID]
	mutex.Unlock()
	if exists {
		wsMessage := models.WebSocketMessage{
			Type:    "sendmessage",
			From:    userID,
			To:      recipientID,
			Content: message.Content,
		}
		msgBytes, err := json.Marshal(wsMessage)
		if err != nil {
			fmt.Printf("Failed to marshal WebSocket message: %v\n", err)
		} else {
			err = recipientConn.WriteMessage(websocket.TextMessage, msgBytes)
			if err != nil {
				fmt.Printf("Failed to send WebSocket message to recipient %s: %v\n", recipientID, err)
			} else {
				fmt.Printf("Message sent to recipient %s via WebSocket\n", recipientID)
			}
		}
	} else {
		fmt.Printf("Recipient %s is not connected via WebSocket\n", recipientID)
	}
	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully", "id": userID})
}
