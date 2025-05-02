package main

import (
	"log"

	authserviceimpl "github.com/AliMumtaz001/Go_Chat_App/api/auth_service"
	userserviceimpl "github.com/AliMumtaz001/Go_Chat_App/api/user_service"
	"github.com/AliMumtaz001/Go_Chat_App/database"
	"github.com/AliMumtaz001/Go_Chat_App/database/mongodb"
	"github.com/AliMumtaz001/Go_Chat_App/database/postgresdb"
	"github.com/AliMumtaz001/Go_Chat_App/routes"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic("Error loading .env file: %s", err)
	}
	conn, err := postgresdb.DbConnection()
	connmongo, err := mongodb.MOngoConn()
	userdb := database.NewStorage(conn, connmongo)
	authService := authserviceimpl.NewAuthService(authserviceimpl.NewAuthServiceImpl{
		UserAuth: userdb,
	})
	userService := userserviceimpl.NewUserService(userdb)
	router := routes.NewRouter(authService, userService)
	router.Engine.Run(":8002")
}
