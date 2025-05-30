package websocket_impl

import (
	"github.com/AliMumtazDev/Go_Chat_App/database/mongodb"
	"github.com/AliMumtazDev/socket/web_socket"
	"github.com/gorilla/websocket"
	// connection "github.com/AliMumtazDev/socket/connection"
)

type WebSocketServiceImpl struct {
	Clients map[int]*websocket.Conn
	MongoDB mongodb.Storage
}

func NewWebSocketService(input mongodb.Storage) web_socket.WebSocketService {
	return &WebSocketServiceImpl{
		MongoDB: input,
		Clients: make(map[int]*websocket.Conn),
		// MongoDB: ip.Mongo,
	}
}

type NewWebSocketServiceImpl struct {
	MongoDB mongodb.Storage
}

var _ web_socket.WebSocketService = &WebSocketServiceImpl{}
