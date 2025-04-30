package routes

import (
	authservice "github.com/AliMumtaz001/Go_Chat_App/service/auth_service"
	userservice "github.com/AliMumtaz001/Go_Chat_App/service/user_service"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine      *gin.Engine
	AuthService authservice.AuthService
	UserService userservice.UserService
}

func NewRouter(authService authservice.AuthService, userService userservice.UserService) *Router {
	engine := gin.Default()
	router := &Router{
		Engine:      engine,
		AuthService: authService,
		UserService: userService,
	}
	router.defineRoutes()
	return router
}

