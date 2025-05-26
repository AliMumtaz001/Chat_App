package connection

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var Conn *websocket.Conn
var ConnMutex sync.Mutex

func ConnectToWebSocketServer(url string, token string) {
	ConnMutex.Lock()
	defer ConnMutex.Unlock()

	var err error
	header := http.Header{}
	header.Add("Authorization", "Bearer "+token)
	Conn, _, err = websocket.DefaultDialer.Dial(url, header)
	if err != nil {
		log.Printf("Failed to connect to WebSocket server: %v", err)
		return
	}
	log.Println("Connected to WebSocket server")

	go func() {
		for {
			if _, _, err := Conn.ReadMessage(); err != nil {
				log.Printf("WebSocket read error: %v, attempting reconnect", err)
				ConnMutex.Lock()
				Conn.Close()
				Conn, _, _ = websocket.DefaultDialer.Dial(url, header)
				ConnMutex.Unlock()
				if err != nil {
					log.Printf("Reconnect failed: %v", err)
					return
				}
				log.Println("Reconnected to WebSocket server")
			}
		}
	}()
}