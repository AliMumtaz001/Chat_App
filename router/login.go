package routes

import (
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
