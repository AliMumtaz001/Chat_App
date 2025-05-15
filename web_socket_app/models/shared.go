package models

import (
	"net/http"
	"sync"

	socketimpl "github.com/AliMumtazDev/socket/websocket_impl"
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
	Connections = make(map[string]*socketimpl.Client)
	ConnMutex   = &sync.Mutex{}
)
