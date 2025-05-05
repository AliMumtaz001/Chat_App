package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("secret-key")
var refreshSecretKey = []byte("my_refresh_secret_key")

func CreateToken(email string, id int) (string, error) {
	fmt.Println("User_id", id)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email":   email,
			"user_id": id,
			"exp":     time.Now().Add(time.Hour * 3).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}


func VerifyToken(tokenString string, c *gin.Context) (*jwt.Token, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}


	return token, nil
}


func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		// tokenString = strings.TrimPrefix(tokenString, "Bearer")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}
		splitToken := strings.Split(authHeader, " ")

		if len(splitToken) != 2 || splitToken[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}
		tokenString := splitToken[1]

		// Verify the token
		token, err := VerifyToken(tokenString, c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		// Extract claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			fmt.Println("Invalid token claims 87")
			c.BindJSON(http.StatusBadRequest)
			return
		}
		userID := fmt.Sprintf("%v", claims["user_id"])
		fmt.Println("User ID from token:", userID)
		c.Set("userID", userID)

		c.Next()
	}
}


