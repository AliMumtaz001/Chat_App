package socket

import (
	"github.com/AliMumtazDev/Go_Chat_App/database/mongodb"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn   *websocket.Conn
	UserID string
}

type WebSocketServiceImpl struct {
	clients map[string]*Client
	MongoDB mongodb.Storage
	wsImpl  *WebSocketImpl 
}

func NewWebSocketService(ws mongodb.Storage) WebSocketService {
	return &WebSocketServiceImpl{
		clients: make(map[string]*Client), 
		MongoDB: ws,
		wsImpl:  NewWebSocketImpl(), 
	}
}

func (impl *WebSocketServiceImpl) SendMessage(client *Client, message []byte) error {
	return impl.wsImpl.SendMessage(client, message)
}

func (impl *WebSocketServiceImpl) ReceiveMessage(client *Client) ([]byte, error) {
	return impl.wsImpl.ReceiveMessage(client)
}

func (impl *WebSocketServiceImpl) RegisterWebSocketRoute(c *gin.Context) {
	impl.wsImpl.RegisterWebSocketRoute(c)
}

var _ WebSocketService = &WebSocketServiceImpl{}
