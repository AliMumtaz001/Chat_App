package main

import (
	"log"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	connMongo, err := mongodb.MongoConn()
	if err != nil {
		log.Fatalf("MongoDB connection error: %s", err)
	}
	messagedb := mongodb.NewMongoDb(connMongo)

	websockets := websocketsimpl.NewWebSockets(messagedb)
	messageService := messageserviceimpl.NewMessageService(messagedb, websockets)

	webSocketRouter := router.NewRouter(messageService, websockets)
	err = webSocketRouter.Engine.Run(":8004")
	if err != nil {
		log.Fatalf("Websocket server failed to start: %s", err)
	}

}
