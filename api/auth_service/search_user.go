package authserviceimpl

import (
	"errors"
	"fmt"

	"github.com/AliMumtazDev/Go_Chat_App/models"
	"github.com/gin-gonic/gin"
)

func (s *AuthServiceImpl) SearchUserservice(ctx *gin.Context, username string) ([]models.SearchUser, error) {
	if username == "" {
		return nil, errors.New("username cannot be empty")
	}
	users, err := s.userAuth.SearchUserdb(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("failed to search user: %w", err)
	}
	return users, nil
}
