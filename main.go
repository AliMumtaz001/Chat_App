package main

// @title Go Chat App API
// @version 1.0
// @description This is a chat application API built with Go and Gin.
// @host localhost:8002
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
import (
	"log"

	authserviceimpl "github.com/AliMumtazDev/Go_Chat_App/api/auth_service"
	userserviceimpl "github.com/AliMumtazDev/Go_Chat_App/api/message_service"
	"github.com/AliMumtazDev/Go_Chat_App/database/mongodb"
	"github.com/AliMumtazDev/Go_Chat_App/database/postgresdb"
	routes "github.com/AliMumtazDev/Go_Chat_App/router"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic("Error loading .env file: %s", err)
	}
	connpostgres, err := postgresdb.PostgresConn()
	connmongo, err := mongodb.MOngoConn()
	// userdb := database.NewStorage(conn, connmongo)
	userdb := postgresdb.NewStorage(connpostgres)
	messagedb := mongodb.NewStorage(connmongo)
	authService := authserviceimpl.NewAuthService(authserviceimpl.NewAuthServiceImpl{
		UserAuth: userdb,
	})
	// userService := userserviceimpl.NewUserService(userserviceimpl.NewUserServiceImpl{
	// 	MessageAuth: messagedb,
	// })
	userService := userserviceimpl.NewUserService(messagedb)
	router := routes.NewRouter(authService, userService)
	router.Engine.Run(":8002")
}
