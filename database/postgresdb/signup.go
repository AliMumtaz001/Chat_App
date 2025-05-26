package postgresdb

import (
	"database/sql"
	"net/http"

	"github.com/AliMumtazDev/Go_Chat_App/models"
	"github.com/gin-gonic/gin"
)

func (u *StorageImpl) SignUpdb(c *gin.Context, req *models.User) *models.User {

	var existingEmail string
	err := u.db.QueryRow("SELECT email FROM employeedata WHERE email = $1", req.Email).Scan(&existingEmail)
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
