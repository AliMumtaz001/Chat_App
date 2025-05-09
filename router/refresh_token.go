package routes

import (
	"github.com/gin-gonic/gin"
)

// RefreshKeyreq godoc
// @Summary      Refresh access token
// @Description  Generate a new access token using refresh token
// @Tags         auth
// @Produce      json
// @Success      200
// @Failure      401
// @Router       /refresh [get]
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
