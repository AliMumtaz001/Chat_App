package routes

import (
	"github.com/gin-gonic/gin"
)

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
