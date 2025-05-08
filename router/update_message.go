package routes

import (
	"github.com/AliMumtaz001/Go_Chat_App/models"
	"github.com/gin-gonic/gin"
)

// UpdateMessagereq godoc
// @Summary      Update a message
// @Description  Update the content of an existing message
// @Tags         messages
// @Accept       json
// @Produce      json
// @Param        _id      path      string         true  "Message ID"
// @Param        message  body      models.Message  true  "Updated message data"
// @Success      200
// @Failure      400
// @Failure      401
// @Failure      500
// @Security     BearerAuth
// @Router       /update-message/{_id} [put]
func (r *Router) UpdateMessagereq(c *gin.Context) {
	messageID := c.Param("_id")
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(401, gin.H{"error": "User not authenticated"})
		return
	}
	var message models.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	if messageID == "" {
		c.JSON(400, gin.H{"error": "Missing message ID"})
		return
	}
	if message.Content == "" {
		c.JSON(400, gin.H{"error": "Missing new message"})
		return
	}
	err := r.UserService.UpdateMessageservice(c, messageID, message, userID.(string))
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error: " + err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Message updated successfully"})
}
