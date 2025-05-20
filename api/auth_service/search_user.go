package authserviceimpl

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (s *AuthServiceImpl) SearchUserservice(ctx *gin.Context, username string) (bool, error) {

	if username == "" {
		return false, errors.New("query cannot be empty")
	}

	exists, err := s.userAuth.SearchUserdb(ctx, username)
	if err != nil {
		return false, fmt.Errorf("failed to search user: %w", err)
	}

	return exists, nil
}
