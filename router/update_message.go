package routes

import (
	"github.com/AliMumtaz001/Go_Chat_App/models"
	"github.com/gin-gonic/gin"
)

func (r *Router) UpdateMessagereq(c *gin.Context) {
    messageID := c.Param("_id")
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

    err := r.UserService.UpdateMessageservice(c, messageID, message)
    if err != nil {
        c.JSON(500, gin.H{"error": "Internal server error: " + err.Error()})
        return
    }

    c.JSON(200, gin.H{"message": "Message updated successfully"})
}
