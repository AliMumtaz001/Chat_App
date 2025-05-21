package routes

import (
	userserviceimpl "github.com/AliMumtazDev/Go_Chat_App/api/message_service"
	"github.com/gin-contrib/cors"

	socketinterface "github.com/AliMumtazDev/socket/web_socket"
	"github.com/gin-gonic/gin"
)

type SocketRouter struct {
	Engine         *gin.Engine
	WebSocket      socketinterface.WebSocketService
	Messageservice userserviceimpl.UserService
}

func NewRouter(userService userserviceimpl.UserService, websocket socketinterface.WebSocketService) *SocketRouter {
	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowWebSockets:  true,
	}))
	router := &SocketRouter{
		Engine:         engine,
		WebSocket:      websocket,
		Messageservice: userService,
	}
	router.SocketRoutes()
	return router
}
