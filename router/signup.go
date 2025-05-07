package routes

import (
	"net/http"

	"github.com/AliMumtaz001/Go_Chat_App/models"
	"github.com/gin-gonic/gin"
)

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
