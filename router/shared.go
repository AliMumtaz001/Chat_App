package routes

import (
	"sync"

	"github.com/gorilla/websocket"
)

var connections = make(map[string]*websocket.Conn)
var mutex = &sync.Mutex{}
