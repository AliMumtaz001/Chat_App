package userserviceimpl

import (
	"github.com/AliMumtaz001/Go_Chat_App/database"
	userservice "github.com/AliMumtaz001/Go_Chat_App/service/user_service"
	"github.com/gin-gonic/gin"
	"github.com/AliMumtaz001/Go_Chat_App/models"
)

type UserServiceImpl struct {
	UserAuth database.Storage
}


func (u *UserServiceImpl) SearchUsers(ctx *gin.Context, query string) ([]models.User, error) {
	// Implement the logic for searching users here.
	// For now, return a placeholder response.
	return []models.User{}, nil
}

func NewUserService(input database.Storage) userservice.UserService {
	return &UserServiceImpl{
		UserAuth: input,
	}
}

type NewUserServiceImpl struct {
	UserAuth database.Storage
}
