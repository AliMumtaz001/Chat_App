package routes

import (
	userserviceimpl "github.com/AliMumtazDev/Go_Chat_App/api/message_service"

	socketinterface "github.com/AliMumtazDev/socket/web_socket"
	"github.com/gin-gonic/gin"
)

type SocketRouter struct {
	Engine         *gin.Engine
	WebSocket      socketinterface.WebSocketService
	Messageservice userserviceimpl.UserService
}
func NewRouter(userService userserviceimpl.UserService, websocket socketinterface.WebSocketService, onlyWS bool) *SocketRouter {
	engine := gin.Default()
	router := &SocketRouter{
		Engine:         engine,
		WebSocket:      websocket,
		Messageservice: userService,
	}
	router.SocketRoutes()
	return router
}
