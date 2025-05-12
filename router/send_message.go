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
