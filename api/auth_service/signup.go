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
		c.JSON(400, gin.H{"error": "Invalid email format"})
		return nil
	}

	if len(req.Password) < 8 {
		c.JSON(400, gin.H{"error": "Password must be at least 8 characters long"})
		return nil
	}
	hasLetter := false
	hasDigit := false
	for _, c := range req.Password {
		switch {
		case 'a' <= c && c <= 'z', 'A' <= c && c <= 'Z':
			hasLetter = true
		case '0' <= c && c <= '9':
			hasDigit = true
		}
	}
	if !hasLetter || !hasDigit {
		c.JSON(400, gin.H{"error": "Password must contain both letters and numbers"})
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
