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
	"fmt"
	"log"
	"os"
	authserviceimpl "github.com/AliMumtazDev/Go_Chat_App/api/auth_service"
	userserviceimpl "github.com/AliMumtazDev/Go_Chat_App/api/message_service"
	"github.com/AliMumtazDev/Go_Chat_App/database/mongodb"
	"github.com/AliMumtazDev/Go_Chat_App/database/postgresdb"
	routes "github.com/AliMumtazDev/Go_Chat_App/router"
	connection "github.com/AliMumtazDev/Go_Chat_App/socket_clint"
	"github.com/AliMumtazDev/socket/web_socket/websocket_impl"
	"github.com/gorilla/websocket"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")
	key := os.Getenv("BACKEND_WS_KEY")
	if err != nil {
		log.Panicf("Error loading .env file: %s", err)
	}
	fmt.Println("Using WebSocket key:", key)
	if key != "676767" {
		log.Printf("Invalid key: %s", key)
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
	websocketServiceConfig := websocket_impl.NewWebSocketServiceImpl{
		Clients: make(map[int]*websocket.Conn),
		MongoDB: messagedb,
	}
	websockets := websocket_impl.NewWebSocketService(websocketServiceConfig)
	
	userService := userserviceimpl.NewUserService(messagedb, websockets)
	go func() {
        connection.ConnectToWebSocketServer("ws://localhost:8004/backend/ws", key)
            log.Println("Connected to WebSocket server")
    }()
	httpRouter := routes.NewRouter(authService, userService)
	if err := httpRouter.Engine.Run(":8005"); err != nil {
		log.Fatalf("HTTP server failed to start: %s", err)
	}
}
