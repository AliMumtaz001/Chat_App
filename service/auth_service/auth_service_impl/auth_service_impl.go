package authserviceimpl

import (
	"github.com/AliMumtaz001/Go_Chat_App/database"
	authservice "github.com/AliMumtaz001/Go_Chat_App/service/auth_service"
)

type AuthServiceImpl struct {
	userAuth database.Storage
}

func NewAuthService(input NewAuthServiceImpl) authservice.AuthService {
	return &AuthServiceImpl{
		userAuth: input.UserAuth,
	}
}

type NewAuthServiceImpl struct {
	UserAuth database.Storage
}
