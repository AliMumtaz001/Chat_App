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
	socket "github.com/AliMumtazDev/Go_Chat_App/web_socket"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Panicf("Error loading .env file: %s", err)
	}

	// Initialize database connections
	connpostgres, err := postgresdb.PostgresConn()
	if err != nil {
		log.Panicf("Error connecting to PostgreSQL: %s", err)
	}
	connmongo, err := mongodb.MOngoConn()
	if err != nil {
		log.Panicf("Error connecting to MongoDB: %s", err)
	}

	// Initialize services
	userdb := postgresdb.NewStorage(connpostgres)
	messagedb := mongodb.NewStorage(connmongo)
	authService := authserviceimpl.NewAuthService(authserviceimpl.NewAuthServiceImpl{
		UserAuth: userdb,
	})
	userService := userserviceimpl.NewUserService(messagedb)
	// Create an instance of WebSocketServiceImpl
	webSocketImpl := socket.WebSocketServiceImpl{}
	websocket := socket.NewWebSocketService(webSocketImpl)
	// Initialize Gin router
	router := routes.NewRouter(authService, userService, websocket, true)

	// Start the server
	log.Println("Server is running on port 8080")
	log.Fatal(router.Engine.Run(":8080")) // Use Gin's router to start the server
}
