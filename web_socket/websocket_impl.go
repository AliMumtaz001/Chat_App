package socket

import (
	"github.com/AliMumtazDev/Go_Chat_App/database/mongodb"
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn   *websocket.Conn
	UserID string
}

type WebSocketServiceImpl struct {
	clients map[string]*Client
	MongoDB mongodb.Storage
}

func NewWebSocketService(ws socket.WebSocketServiceImpl) socket.WebSocketService {
	return &WebSocketServiceImpl{
		clients: ws.clients,
		MongoDB: ws.WebSocket,
	}
}

// var _ WebSocketService = &WebSocketServiceImpl{}
