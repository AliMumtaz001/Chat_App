package userserviceimpl

import (
	"github.com/AliMumtaz001/Go_Chat_App/database"
	"github.com/AliMumtaz001/Go_Chat_App/models"
	"github.com/gin-gonic/gin"
)

type UserServiceImpl struct {
	UserAuth database.Storage
}

func (u *UserServiceImpl) SearchUsers(ctx *gin.Context, query string) ([]models.User, error) {
	return []models.User{}, nil
}

func NewUserService(input database.Storage) UserService {
	return &UserServiceImpl{
		UserAuth: input,
	}
}

type NewUserServiceImpl struct {
	UserAuth database.Storage
}
