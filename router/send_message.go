package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/AliMumtazDev/Go_Chat_App/models"
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

	fmt.Println("id", userID)

	// Bind JSON request to the message model
	if err := c.ShouldBindJSON(&message); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	fmt.Printf("Received message: %+v\n", message)

	senderID, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		fmt.Println("Error parsing sender ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sender ID"})
		return
	}

	if err := r.UserService.SendMessageservice(c, &message); err != nil {
		fmt.Println("Error saving message:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message sent from ", "id": senderID, "successfull id": message.ReceiverID})
}
