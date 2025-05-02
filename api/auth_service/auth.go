package authserviceimpl

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/AliMumtaz001/Go_Chat_App/auth"
	"github.com/AliMumtaz001/Go_Chat_App/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var refreshSecretKey = []byte("my_refresh_secret_key")
var secretKey = []byte("secret-key")

func (u *AuthServiceImpl) SignUpservice(c *gin.Context, req *models.User) *models.User {

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

func (u *AuthServiceImpl) Loginservice(c *gin.Context, req *models.UserLogin) (*models.TokenPair, error) {

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

func (a *AuthServiceImpl) RefreshAccessTokenservice(c *gin.Context) (string, error) {
	refreshTokenString := c.GetHeader("Authorization")
	refreshTokenString = strings.TrimPrefix(refreshTokenString, "Bearer ")

	if refreshTokenString == "" {
		return "", fmt.Errorf("Refresh token is empty")
	}

	refreshToken, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(refreshSecretKey), nil
	})
	if err != nil || !refreshToken.Valid {
		return "", fmt.Errorf("Invalid refresh token")
	}

	claims, ok := refreshToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("Invalid refresh token claims")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return "", fmt.Errorf("Invalid email in refresh token")
	}

	newAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(15 * time.Minute).Unix(),
	})

	newAccessTokenString, err := newAccessToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("Failed to generate new access token")
	}

	return newAccessTokenString, nil
}

func (s *AuthServiceImpl) SearchUserservice(ctx *gin.Context, query string) (bool, error) {
	// Basic validation
	if query == "" {
		return false, errors.New("query cannot be empty")
	}

	// Call the repository to check if the user exists
	exists, err := s.userAuth.SearchUserdb(ctx, query)
	if err != nil {
		return false, fmt.Errorf("failed to search user: %w", err)
	}

	return exists, nil
}

func (s *AuthServiceImpl)SendMessageservice(c *gin.Context, sID string, message models.Message) error{
	err := s.userAuth.SendMessagedb(c,sID, message)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}
	return nil
}
