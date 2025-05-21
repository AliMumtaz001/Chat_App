package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SearchUserreq godoc
// @Summary      Search for users with partial matching
// @Description  Returns a list of users whose usernames partially match the query
// @Tags         users
// @Produce      json
// @Param        q  query     string  true  "Search query for username (partial match)"
// @Success      200  {object}  gin.H  "List of matching users"
// @Failure      400  {object}  gin.H  "Query parameter 'user' is required"
// @Failure      500  {object}  gin.H  "Failed to search users"
// @Router       /search-user [get]
func (r *Router) SearchUserreq(c *gin.Context) {
	query := c.Query("user") 
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'user' is required"})
		return
	}
	users, err := r.AuthService.SearchUserservice(c, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}