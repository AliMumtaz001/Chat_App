package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) SearchUserreq(c *gin.Context) {
	username := c.Query("username")
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
