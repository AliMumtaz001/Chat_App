package authserviceimpl

import (
	"github.com/AliMumtazDev/Go_Chat_App/models"
	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Loginservice(c *gin.Context, login *models.UserLogin) (*models.TokenPair, error)
	SignUpservice(c *gin.Context, req *models.User) *models.User
	RefreshAccessTokenservice(c *gin.Context) (string, error)
	SearchUserservice(ctx *gin.Context, username string) ([]models.SearchUser, error) 
}
