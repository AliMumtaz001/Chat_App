package postgresdb

import (
    "fmt"

    "github.com/AliMumtazDev/Go_Chat_App/models"
    "github.com/gin-gonic/gin"
)

func (u *StorageImpl) SearchUserdb(ctx *gin.Context, username string) ([]models.SearchUser, error) {
    var users []models.SearchUser
    querySQL := `SELECT id, username FROM employeedata WHERE username ILIKE $1 LIMIT 10`
    rows, err := u.db.QueryContext(ctx, querySQL, "%"+username+"%")
    if err != nil {
        return nil, fmt.Errorf("failed to query users: %w", err)
    }
    defer rows.Close()
    for rows.Next() {
        var user models.SearchUser
        err := rows.Scan(&user.Id, &user.Username)
        if err != nil {
            return nil, fmt.Errorf("failed to scan user: %w", err)
        }
        users = append(users, user)
    }
    return users, nil
}