package models

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var (
	Connections = make(map[string]*Client)
	ConnMutex   = &sync.Mutex{}
)

type Client struct {
	Conn   *websocket.Conn
	UserID string
}
