package routes

import (
	"net/http"

	"github.com/AliMumtaz001/Go_Chat_App/models"
	"github.com/gin-gonic/gin"
)

func (r *Router) Login(c *gin.Context) {
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

	tokenPair, err := r.AuthService.Login(c, &login)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tokenPair)
}
func (r *Router) SignUp(c *gin.Context) {
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
	response := r.AuthService.SignUp(c, &signup)
	if response == nil {
		return
	}
	c.JSON(http.StatusOK, response)

}

func (r *Router) RefreshKey(c *gin.Context) {
	newToken, err := r.AuthService.RefreshAccessToken(c)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"access_token": newToken,
	})
}

func (r *Router) SearchUserHandler(c *gin.Context) {
	// Get query parameter (e.g., email or username)
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'q' is required"})
		return
	}

	// Call the auth service to search for the user
	exists, err := r.AuthService.SearchUser(c, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search user"})
		return
	}

	// Return response
	if exists {
		c.JSON(http.StatusOK, gin.H{"message": "User exists"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
	}
}
