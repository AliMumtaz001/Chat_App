package routes

import (
	"github.com/AliMumtazDev/Go_Chat_App/auth"
	"github.com/gin-gonic/gin"
)

func (r *SocketRouter) SocketRoutes() {
	r.Engine.GET("/ws", auth.AuthMiddleware(), func(c *gin.Context) {
		r.WebSocket.RegisterWebSocketRoute(c)
	})
}
