package main

import (
	"log"
	"github.com/AliMumtaz001/Go_Chat_App/database"
	"github.com/AliMumtaz001/Go_Chat_App/routes"
	authserviceimpl "github.com/AliMumtaz001/Go_Chat_App/service/auth_service/auth_service_impl"
	userserviceimpl "github.com/AliMumtaz001/Go_Chat_App/service/user_service/user_service_impl"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	conn, err := database.DbConnection()
	userdb := database.NewStorage(conn)
	authService := authserviceimpl.NewAuthService(authserviceimpl.NewAuthServiceImpl{
		UserAuth: userdb,
	})
	userService := userserviceimpl.NewUserService(userdb)
	router := routes.NewRouter(authService, userService)
	router.Engine.Run(":8002")

}
