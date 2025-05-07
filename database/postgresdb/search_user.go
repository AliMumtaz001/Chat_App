package postgresdb

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (r *StorageImpl) SearchUserdb(c *gin.Context, query string) (bool, error) {
	var exists bool
	querySQL := `SELECT EXISTS (SELECT 1 FROM employeedata WHERE username = $1)`
	err := r.db.QueryRowContext(c, querySQL, query).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to query user: %w", err)
	}
	return exists, nil
}
