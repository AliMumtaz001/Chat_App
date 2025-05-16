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
	connection "github.com/AliMumtazDev/Go_Chat_App/socket_clint"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panicf("Error loading .env file: %s", err)
	}

	connpostgres, err := postgresdb.PostgresConn()
	if err != nil {
		log.Panicf("Error connecting to PostgreSQL: %s", err)
	}
	connmongo, err := mongodb.MOngoConn()
	if err != nil {
		log.Panicf("Error connecting to MongoDB: %s", err)
	}

	userdb := postgresdb.NewStorage(connpostgres)
	messagedb := mongodb.NewStorage(connmongo)
	authService := authserviceimpl.NewAuthService(authserviceimpl.NewAuthServiceImpl{
		UserAuth: userdb,
	})
	userService := userserviceimpl.NewUserService(messagedb)
	// Create an instance of WebSocketServiceImpl
	// webSocketImpl := sock et.WebSocketServiceImpl{}
	// websocket := websocket_impl.NewWebSocketService(messagedb)
	// router := routes.NewRouter(authService, userService, websocket, true)
	// log.Println("Server is running on port 8005")
	//16

	stoken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXIxQGdtYWlsLmNvbSIsImV4cCI6MTc0NzQwOTYzMywidXNlcl9pZCI6MTd9.ldN_Q3PgGatgbfFYxDgE2RoDXwUXbzRBQQhI00Zp2o0"
	//15
	rtoken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXIyQGdtYWlsLmNvbSIsImV4cCI6MTc0NzQwOTY2NSwidXNlcl9pZCI6MTh9.vwsvDpezybda97yOIrG7ouUyHhu6DbrgxU_ID8ui6Mw"
	go connection.ConnectToWebSocketServer("ws://localhost:8003/ws", stoken)
	go connection.ConnectToWebSocketServer("ws://localhost:8003/ws", rtoken)
	httpRouter := routes.NewRouter(authService, userService, false)
	if err := httpRouter.Engine.Run(":8005"); err != nil {
		log.Fatalf("HTTP server failed to start: %s", err)
	}

}
