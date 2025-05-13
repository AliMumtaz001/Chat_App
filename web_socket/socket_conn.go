package socket

import (
	"log"
	"net/http"
)

type WebSocketImpl struct{}

func NewWebSocketImpl() *WebSocketImpl {
	return &WebSocketImpl{}
}

func (ws *WebSocketImpl) UpgradeConnection(w http.ResponseWriter, r *http.Request) (*Client, error) {
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %v", err)
		return nil, err
	}

	userID := r.URL.Query().Get("userID") 

	client := &Client{
		Conn:   conn,
		UserID: userID,
	}

	ConnMutex.Lock()
	Connections[userID] = client
	ConnMutex.Unlock()

	return client, nil
}
