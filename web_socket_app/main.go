package main

import (
	"log"

	userserviceimpl "github.com/AliMumtazDev/Go_Chat_App/api/message_service"
	"github.com/AliMumtazDev/Go_Chat_App/database/mongodb"
	routes "github.com/AliMumtazDev/socket/router"
	"github.com/AliMumtazDev/socket/web_socket/websocket_impl"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	connMongo, err := mongodb.MOngoConn()
	if err != nil {
		log.Fatalf("MongoDB connection error: %s", err)
	}
	messagedb := mongodb.NewStorage(connMongo)
	
	websockets := websocket_impl.NewWebSocketService(messagedb)
	messageService := userserviceimpl.NewUserService(messagedb, websockets)
	webSocketRouter := routes.NewRouter(messageService, websockets)
	err = webSocketRouter.Engine.Run(":8004")
	if err != nil {
		log.Fatalf("Websocket server failed to start: %s", err)
	}
}
