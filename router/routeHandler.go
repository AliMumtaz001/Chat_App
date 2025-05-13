package routes

import (
	"github.com/AliMumtazDev/Go_Chat_App/auth"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (r *Router) defineRoutes() {
	r.Engine.POST("/signup", r.SignUpreq)
	r.Engine.POST("/login", r.Loginreq)
	r.Engine.GET("/refresh", r.RefreshKeyreq)
	r.Engine.GET("/search-user", r.SearchUserreq)
	r.Engine.POST("/sendmessage", auth.AuthMiddleware(), r.SendMessagereq)
	r.Engine.GET("/getmessage", auth.AuthMiddleware(), r.GetMessagereq)
	r.Engine.PUT("/update-message/:_id", auth.AuthMiddleware(), r.UpdateMessagereq)
	r.Engine.POST("/delete-message/:_id", auth.AuthMiddleware(), r.DeleteMessagereq)
	r.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (r *Router) socketRoutes() {
	// WebSocket route for handling WebSocket connections
	r.Engine.GET("/ws", func(c *gin.Context) {
		r.WebSocket.RegisterWebSocketRoute(c)
	})

	// Add more WebSocket-specific routes here if needed
}
