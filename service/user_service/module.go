package userservice

import (
	"github.com/AliMumtaz001/Go_Chat_App/models"
	"github.com/gin-gonic/gin"
)

type UserService interface {
	// SaveResults(c *gin.Context, req *models.Result) *models.Result
	// GetUserResults(c *gin.Context, page int, limit int) ([]models.Result, error)
	SearchUsers(c *gin.Context, query string) ([]models.User, error)
}
