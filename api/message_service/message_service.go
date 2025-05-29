package userserviceimpl

import (
	"github.com/AliMumtazDev/Go_Chat_App/database/mongodb"
	"github.com/AliMumtazDev/socket/web_socket"
)

type UserServiceImpl struct {
	messageAuth mongodb.Storage
	WebSocket   web_socket.WebSocketService
}
func NewUserService(input mongodb.Storage, ws web_socket.WebSocketService) UserService {
	return &UserServiceImpl{
		messageAuth: input,
		WebSocket: ws,
	}
}

type NewUserServiceImpl struct {
	messageAuth mongodb.Storage
}

var _ UserService = &UserServiceImpl{}
