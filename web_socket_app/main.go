package main

import (
	"log"

	userserviceimpl "github.com/AliMumtazDev/Go_Chat_App/api/message_service"
	"github.com/AliMumtazDev/Go_Chat_App/database/mongodb"
	"github.com/AliMumtazDev/socket/client"
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

	webSocketImpl := websocket_impl.WebSocketServiceImpl{
		Clients: make(map[string]*client.Client),
		MongoDB: messagedb,
	}

	websockets := websocket_impl.NewWebSocketService(webSocketImpl)
	messageService := userserviceimpl.NewUserService(messagedb)

	webSocketRouter := routes.NewRouter(messageService, websockets, false)
	err = webSocketRouter.Engine.Run(":8004")
	if err != nil {
		log.Fatalf("Websocket server failed to start: %s", err)
	}

}
