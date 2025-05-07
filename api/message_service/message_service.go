package userserviceimpl

import (
	"github.com/AliMumtaz001/Go_Chat_App/database/mongodb"
)

type UserServiceImpl struct {
	messageAuth mongodb.Storage
}

func NewUserService(input mongodb.Storage) *UserServiceImpl {
	return &UserServiceImpl{
		messageAuth: input,
	}
}

type NewUserServiceImpl struct {
	messageAuth mongodb.Storage
}

var _ UserService = &UserServiceImpl{}
