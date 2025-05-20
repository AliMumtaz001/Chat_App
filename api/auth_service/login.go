package authserviceimpl

import (
	"errors"
	"regexp"
	"time"

	"github.com/AliMumtazDev/Go_Chat_App/auth"
	"github.com/AliMumtazDev/Go_Chat_App/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (u *AuthServiceImpl) Loginservice(c *gin.Context, req *models.UserLogin) (*models.TokenPair, error) {

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched := regexp.MustCompile(emailRegex).MatchString(req.Email)
	if !matched {
		return nil, errors.New("Invalid email format")
	}

	user, err := u.userAuth.FindUserByEmaildb(req.Email)
	if err != nil {
		return nil, errors.New("User not found")
	}

	if user.Password != req.Password {
		return nil, errors.New("Invalid credentials")
	}

	token, err := auth.CreateToken(user.Email, int(user.Id))
	if err != nil {
		return nil, errors.New("Failed to generate access token")
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(7 * 24 * time.Hour).Unix(),
	})

	refreshTokenString, err := refreshToken.SignedString(refreshSecretKey)
	if err != nil {
		return nil, errors.New("Failed to generate refresh token")
	}

	response := models.TokenPair{
		AccessToken:  token,
		RefreshToken: refreshTokenString,
	}
	return &response, nil
}
