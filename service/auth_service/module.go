package authservice

import (
	"github.com/AliMumtaz001/Go_Chat_App/models"
	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Login(c *gin.Context, login *models.UserLogin) (*models.TokenPair, error)
	SignUp(c *gin.Context, req *models.User) *models.User
	RefreshAccessToken(c *gin.Context) (string, error)
	
}
