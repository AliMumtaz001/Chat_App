package auth

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

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

		token, err := VerifyToken(tokenString, c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			fmt.Println("Invalid token claims 87")
			c.BindJSON(http.StatusBadRequest)
			return
		}
		userID := fmt.Sprintf("%v", claims["user_id"])
		c.Set("userID", userID)

		c.Next()
	}
}

func WSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := c.Query("token")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token query parameter missing"})
			c.Abort()
			return
		}
		fmt.Println("Token from query:", tokenString)

		token, err := VerifyToken(tokenString, c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: " + err.Error()})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		userID := fmt.Sprintf("%v", claims["user_id"])
		c.Set("userID", userID)

		c.Next()
	}
}

func BackendWSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := godotenv.Load(".env")
		if err != nil {
			log.Printf("Error loading .env file: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Server configuration error"})
			c.Abort()
			return
		}

		receivedKey := c.Query("key")
		if receivedKey == "" {
			log.Println("Missing key query parameter")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token query parameter missing"})
			c.Abort()
			return
		}

		key := os.Getenv("BACKEND_WS_KEY")
		if key != receivedKey {
			log.Printf("Invalid key: received=%s, expected=%s", receivedKey, key)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid key"})
			c.Abort()
			return
		}

		log.Println("Backend WebSocket key validated successfully")
		c.Set("userID", "-1")
		c.Next()
	}
}
