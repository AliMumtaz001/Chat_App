package postgresdb

import (
	"database/sql"
	"fmt"

	"github.com/AliMumtazDev/Go_Chat_App/models"
)

type StorageImpl struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &StorageImpl{
		db: db,
	}
}

type UserService struct {
	Db *sql.DB
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
