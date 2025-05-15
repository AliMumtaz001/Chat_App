package routes

import (
	socketinterface "github.com/AliMumtazDev/socket/web_socket"
	"github.com/gin-gonic/gin"
)

type SocketRouter struct {
	Engine *gin.Engine
	// WebSocket socketinterface.WebSocketService
	WebSocket   socketinterface.WebSocketService
	AuthService authservice.AuthService
}

// func NewRouter(authService authservice.AuthService, userService userservice.UserService, websocket socketinterface.WebSocketService, onlyWS bool) *Router {
func NewRouter(authService authService.AuthService, userService userservice.UserService, websocket socketinterface.WebSocketService, onlyWS bool) *SocketRouter {
	engine := gin.Default()
	router := &SocketRouter{
		Engine:    engine,
		WebSocket: websocket,
	}
	router.SocketRoutes()
	return router
}
