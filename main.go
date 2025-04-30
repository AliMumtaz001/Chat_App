package main

import (
	"log"

	"github.com/AliMumtaz001/Go_Chat_App/database"
	"github.com/AliMumtaz001/Go_Chat_App/routes"
	authserviceimpl "github.com/AliMumtaz001/Go_Chat_App/service/auth_service"
	userserviceimpl "github.com/AliMumtaz001/Go_Chat_App/service/user_service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic("Error loading .env file: %s", err)
	}
	conn, err := database.DbConnection()
	connmongo, err := database.MOngoConn()
	userdb := database.NewStorage(conn, connmongo)
	authService := authserviceimpl.NewAuthService(authserviceimpl.NewAuthServiceImpl{
		UserAuth: userdb,
	})
	userService := userserviceimpl.NewUserService(userdb)
	router := routes.NewRouter(authService, userService)
	router.Engine.Run(":8002")
}
