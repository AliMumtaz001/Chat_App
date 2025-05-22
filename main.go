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
	// connection "github.com/AliMumtazDev/Go_Chat_App/socket_clint"
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
	//17

// 	stoken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXIxQGdtYWlsLmNvbSIsImV4cCI6MTc0NzY0MTM3MiwidXNlcl9pZCI6MTd9.3semoQBPir4Nw7iit94gIQPQzDNcN-Lj-KX04OmpDQs"
// 	//18
// 	rtoken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXIyQGdtYWlsLmNvbSIsImV4cCI6MTc0NzY0MTQwOSwidXNlcl9pZCI6MTh9.nqTCEI8of7vlRJA9CaFADB0wm4f5BYinljwP1ytwF7I"
// 	// ...existing code...
// go connection.ConnectToWebSocketServer("ws://localhost:8004/protected/ws", stoken)
// go connection.ConnectToWebSocketServer("ws://localhost:8004/protected/ws", rtoken)
// // ...existing code...
	httpRouter := routes.NewRouter(authService, userService, false)
	if err := httpRouter.Engine.Run(":8005"); err != nil {
		log.Fatalf("HTTP server failed to start: %s", err)
	}

}
