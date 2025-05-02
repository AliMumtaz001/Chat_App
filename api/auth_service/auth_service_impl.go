package authserviceimpl

import (
	"github.com/AliMumtaz001/Go_Chat_App/database"
)

type AuthServiceImpl struct {
	userAuth database.Storage
}

func NewAuthService(input NewAuthServiceImpl) AuthService {
	return &AuthServiceImpl{
		userAuth: input.UserAuth,
	}
}

type NewAuthServiceImpl struct {
	UserAuth database.Storage
}
