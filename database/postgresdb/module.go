package postgresdb

import (
	"github.com/AliMumtaz001/Go_Chat_App/models"
	"github.com/gin-gonic/gin"
)

type Storage interface {
	FindUserByEmaildb(email string) (*models.UserLogin, error)
	SignUpdb(c *gin.Context, req *models.User) *models.User
	SearchUserdb(ctx *gin.Context, username string) (bool, error)
}
