package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/AliMumtaz001/Go_Chat_App/models"
	"github.com/gin-gonic/gin"
)

func (r *Router) Loginreq(c *gin.Context) {
	var req models.UserLoginReq
	var login models.UserLogin

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	login = models.UserLogin{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	tokenPair, err := r.AuthService.Loginservice(c, &login)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tokenPair)
}
func (r *Router) SignUpreq(c *gin.Context) {
	var req *models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	signup := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
	response := r.AuthService.SignUpservice(c, &signup)
	if response == nil {
		return
	}
	c.JSON(http.StatusOK, response)

}

func (r *Router) RefreshKeyreq(c *gin.Context) {
	newToken, err := r.AuthService.RefreshAccessTokenservice(c)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"access_token": newToken,
	})
}

func (r *Router) SearchUserreq(c *gin.Context) {
	// Get query parameter (e.g., email or username)
	username := c.Query("q")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'q' is required"})
		return
	}

	exists, err := r.AuthService.SearchUserservice(c, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search user"})
		return
	}

	if exists {
		c.JSON(http.StatusOK, gin.H{"message": "User exists"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
	}
}

func (r *Router) SendMessagereq(c *gin.Context) {
	// Get the message from the request body
	var message models.Message

	userID := c.MustGet("userID").(string)
	if err := c.ShouldBindJSON(&message); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	receiverID := strconv.FormatInt(message.ReceiverID, 10)

	// Call the service layer to send the message
	err := r.UserService.SendMessageservice(c, userID, receiverID, message)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully", "id": userID})
}

func (r *Router) GetMessagereq(c *gin.Context) {
	// var msg models.Message
	// Get query parameters
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

	// Debug log
	fmt.Printf("Sender ID: %s, Receiver ID: %s\n", senderID, receiverID)

	// Call the service layer to get messages
	messages, err := r.UserService.GetMessageservice(c, senderID, receiverID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, messages)
}

// }
