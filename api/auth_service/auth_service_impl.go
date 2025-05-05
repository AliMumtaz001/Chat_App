package authserviceimpl

import (
	"github.com/AliMumtaz001/Go_Chat_App/database/postgresdb"
)

type AuthServiceImpl struct {
	userAuth postgresdb.Storage
}

func NewAuthService(input NewAuthServiceImpl) AuthService {
	return &AuthServiceImpl{
		userAuth: input.UserAuth,
	}
}

type NewAuthServiceImpl struct {
	UserAuth postgresdb.Storage
}

var _ AuthService = &AuthServiceImpl{}
