package authserviceimpl

import (
	"github.com/AliMumtaz001/Go_Chat_App/models"
	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Loginservice(c *gin.Context, login *models.UserLogin) (*models.TokenPair, error)
	SignUpservice(c *gin.Context, req *models.User) *models.User
	RefreshAccessTokenservice(c *gin.Context) (string, error)
	SearchUserservice(c *gin.Context, query string) (bool, error)
	SendMessageservice(c *gin.Context, sID string, message models.Message) error
}
