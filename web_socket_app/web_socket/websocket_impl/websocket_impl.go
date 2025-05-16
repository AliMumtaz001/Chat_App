package websocket_impl

import (
	"github.com/AliMumtazDev/Go_Chat_App/database/mongodb"
	"github.com/AliMumtazDev/socket/client"
	upgradeconn "github.com/AliMumtazDev/socket/connection"
	socketinterface "github.com/AliMumtazDev/socket/web_socket"
)

type WebSocketServiceImpl struct {
	Clients map[string]*client.Client
	MongoDB mongodb.Storage
	WSImpl  *upgradeconn.WebSocketImpl
} //*socket.WebSocketService

var _ socketinterface.WebSocketService = &WebSocketServiceImpl{}

func NewWebSocketService(ws WebSocketServiceImpl) socketinterface.WebSocketService {
	return &WebSocketServiceImpl{
		Clients: make(map[string]*client.Client),
		MongoDB: ws.MongoDB,
		WSImpl:  upgradeconn.NewWebSocketImpl(),
	}
}
