package routes

func (r *Router) defineRoutes() {
	r.Engine.POST("/signup", r.SignUp)
	r.Engine.POST("/login", r.Login)
	r.Engine.GET("/refresh", r.RefreshKey)
	r.Engine.GET("/search", r.SearchUsers)
}
