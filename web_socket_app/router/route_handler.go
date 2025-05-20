package routes
import (
	"github.com/AliMumtazDev/Go_Chat_App/auth"
)
func (r *SocketRouter) SocketRoutes() {
	protected := r.Engine.Group("/protected")
	protected.Use(auth.AuthMiddleware())
	{
		protected.GET("/ws", r.RegisterWebSocketRoute)
	}
}
