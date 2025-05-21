package routes

import (
	authservice "github.com/AliMumtazDev/Go_Chat_App/api/auth_service"
	userservice "github.com/AliMumtazDev/Go_Chat_App/api/message_service"
	"github.com/gin-contrib/cors" 
	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine      *gin.Engine
	AuthService authservice.AuthService
	UserService userservice.UserService
}

func NewRouter(authService authservice.AuthService, userService userservice.UserService, onlyWS bool) *Router {
	engine := gin.Default()
	    engine.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{"GET", "POST", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        AllowWebSockets:  true,
    }))
	router := &Router{
		Engine:      engine,
		AuthService: authService,
		UserService: userService,
	}
	router.defineRoutes()
	return router
}
