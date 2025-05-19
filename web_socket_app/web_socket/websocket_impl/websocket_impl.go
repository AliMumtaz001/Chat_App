package websocket_impl

import (
	"github.com/AliMumtazDev/Go_Chat_App/database/mongodb"
	"github.com/AliMumtazDev/socket/client"

	// connection "github.com/AliMumtazDev/socket/connection"
	socketinterface "github.com/AliMumtazDev/socket/web_socket"
)

type WebSocketServiceImpl struct {
	Clients map[int]*client.Client
	MongoDB mongodb.Storage
	WSImpl  *WebSocketImpl
}

var _ socketinterface.WebSocketService = &WebSocketServiceImpl{}

func NewWebSocketService(ws WebSocketServiceImpl) socketinterface.WebSocketService {
	return &WebSocketServiceImpl{
		Clients: make(map[int]*client.Client),
		MongoDB: ws.MongoDB,
		WSImpl:  &WebSocketImpl{},
	}
}
