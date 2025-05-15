package routes

import "github.com/gin-gonic/gin"

func (r *SocketRouter) SocketRoutes() {
	r.Engine.GET("/ws", func(c *gin.Context) {
		r.WebSocket.RegisterWebSocketRoute(c)
	})
}
