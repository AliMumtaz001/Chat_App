package routes

import (
	"fmt"
	"net/http"

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
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'q' is required"})
		return
	}

	exists, err := r.AuthService.SearchUserservice(c, query)
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
	fmt.Printf("Id here", userID)
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Call the service layer to send the message
	err := r.AuthService.SendMessageservice(c, userID, message)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully", "id": userID})
}
