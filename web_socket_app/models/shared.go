package models

import (
	"net/http"
	"sync"

	"github.com/AliMumtazDev/socket/client"
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
	Connections = make(map[string]*client.Client)
	ConnMutex   = &sync.Mutex{}
)
