package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (r *Router) DeleteMessagereq(c *gin.Context) {
	messageID := c.Param("_id")
	fmt.Println("msg id:", messageID)
	err := r.UserService.DeleteMessageservice(c, messageID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Message deleted successfully"})
}
