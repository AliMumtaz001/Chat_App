package authserviceimpl

import (
	"fmt"
	"regexp"

	"github.com/AliMumtazDev/Go_Chat_App/models"
	"github.com/gin-gonic/gin"
)

func (u *AuthServiceImpl) SignUpservice(c *gin.Context, req *models.User) *models.User {

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched := regexp.MustCompile(emailRegex).MatchString(req.Email)
	if !matched {
		fmt.Println("Invalid email format")
		return nil
	}

	createdUser := u.userAuth.SignUpdb(c, req)

	if createdUser == nil {
		fmt.Println("signup error")
		return nil
	}

	response := models.User{
		Username: createdUser.Username,
		Email:    createdUser.Email,
		Password: createdUser.Password,
		Message:  "User created successfully",
	}

	return &response
}
