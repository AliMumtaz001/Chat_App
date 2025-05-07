package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/AliMumtaz001/Go_Chat_App/models"
	"github.com/gin-gonic/gin"
)

func (r *Router) SendMessagereq(c *gin.Context) {
	var message models.Message

	userID := c.MustGet("userID").(string)
	if err := c.ShouldBindJSON(&message); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	
	sID, _ := strconv.ParseInt(userID, 10, 64)
	err := r.UserService.SendMessageservice(c, sID, message.ReceiverID, message)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully", "id": userID})
}
