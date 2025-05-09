package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// DeleteMessagereq godoc
// @Summary      Delete a message
// @Description  Delete a message by ID
// @Tags         messages
// @Produce      json
// @Param        _id  path      string  true  "Message ID"
// @Success      200
// @Failure      401
// @Failure      500
// @Security     BearerAuth
// @Router       /delete-message/{_id} [post]
func (r *Router) DeleteMessagereq(c *gin.Context) {
	messageID := c.Param("_id")
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(401, gin.H{"error": "User not authenticated"})
		return
	}
	fmt.Println("msg id:", messageID, "userID:", userID)
	err := r.UserService.DeleteMessageservice(c, messageID, userID.(string))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Message deleted successfully"})
}