package routes

import (
	"log"
	"net/http"

	"github.com/AliMumtazDev/Go_Chat_App/models"
	"github.com/gin-gonic/gin"
)

// @Summary      Sign up a new user
// @Description  Create a new user account
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      models.User  true  "User data"
// @Success      200   {object}  map[string]interface{}
// @Failure      400
// @Router       /signup [post]
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

		log.Println("66", response)
		return
	}
	c.JSON(http.StatusOK, response)

}
