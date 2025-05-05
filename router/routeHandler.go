package routes

import (
	"github.com/AliMumtaz001/Go_Chat_App/auth"
)

func (r *Router) defineRoutes() {
	r.Engine.POST("/signup", r.SignUpreq)
	r.Engine.POST("/login", r.Loginreq)
	r.Engine.GET("/refresh", r.RefreshKeyreq)
	r.Engine.GET("/search-user", r.SearchUserreq)
	r.Engine.POST("/sendmessage", auth.AuthMiddleware(), r.SendMessagereq)
	r.Engine.GET("/getmessage", auth.AuthMiddleware(), r.GetMessagereq)
}
