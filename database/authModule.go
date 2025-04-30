package database

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/AliMumtaz001/Go_Chat_App/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type StorageImpl struct {
	db          *sql.DB
	mongoClient *mongo.Client
}

func NewStorage(db *sql.DB, input *mongo.Client) Storage {
	return &StorageImpl{
		db:          db,
		mongoClient: input,
	}
}

func (u *StorageImpl) FindUserByEmaildb(email string) (*models.UserLogin, error) {
	fmt.Println(59)

	var user models.UserLogin

	err := u.db.QueryRow("SELECT id, username, email, password FROM employeedata WHERE email=$1", email).Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	if err != nil {
		fmt.Println(err, "22")
		return nil, err
	}

	return &user, nil

}

func (u *StorageImpl) SignUpdb(c *gin.Context, req *models.User) *models.User {

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, err := regexp.MatchString(emailRegex, req.Email)
	if err != nil || !matched {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return nil
	}

	var existingEmail string
	err = u.db.QueryRow("SELECT email FROM employeedata WHERE email = $1", req.Email).Scan(&existingEmail)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists with this email"})
		return nil
	} else if err != sql.ErrNoRows {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking existing email"})
		return nil
	}

	var existingUsername string
	err = u.db.QueryRow("SELECT username FROM employeedata WHERE username = $1", req.Username).Scan(&existingUsername)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username not available, choose another"})
		return nil
	} else if err != sql.ErrNoRows {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking existing username"})
		return nil
	}

	_, err = u.db.Exec("INSERT INTO employeedata (username, email, password) VALUES ($1, $2, $3)", req.Username, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user", "details": err.Error()})
		return nil
	}

	return req
}

// UserService represents a service with a database connection
type UserService struct {
	Db *sql.DB
}

func (r *StorageImpl) SearchUserdb(c *gin.Context, query string) (bool, error) {
	var exists bool
	querySQL := `SELECT EXISTS (SELECT 1 FROM employeedata WHERE username = $1)`
	err := r.db.QueryRowContext(c, querySQL, query).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to query user: %w", err)
	}
	return exists, nil
}

func (r *StorageImpl) SendMessagedb(c *gin.Context, msg models.Message) error {
	message := models.Message{
		SenderID:   msg.SenderID,
		ReceiverID: msg.ReceiverID,
		Content:    msg.Content,
		Timestamp:  time.Now(),
	}
	collection := r.mongoClient.Database("chatdb").Collection("sendmsg")
	_, err := collection.InsertOne(context.TODO(), message)
	return err
}
