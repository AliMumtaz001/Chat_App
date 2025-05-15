package routes

import (
	authservice "github.com/AliMumtazDev/Go_Chat_App/api/auth_service"
	userservice "github.com/AliMumtazDev/Go_Chat_App/api/message_service"
	socket "github.com/AliMumtazDev/Go_Chat_App/web_socket"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine      *gin.Engine
	AuthService authservice.AuthService
	UserService userservice.UserService
	WebSocket   socket.WebSocketService
}

func NewRouter(authService authservice.AuthService, userService userservice.UserService, websocket socket.WebSocketService, onlyWS bool) *Router {
	engine := gin.Default()
	router := &Router{
		Engine:      engine,
		AuthService: authService,
		UserService: userService,
		WebSocket:   websocket,
	}
	if onlyWS {
		router.socketRoutes()
	} else {
		router.defineRoutes()
	}
	return router
}
