package models

type User struct {
	// gorm.Model
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Message  string `json: "message"`
}

type UserLogin struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type UserLoginReq struct {
	Username string `json:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
