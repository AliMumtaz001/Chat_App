package upgradeconn

import (
	"log"
	"net/http"
	"sync"

	// socket "github.com/AliMumtazDev/socket/models"
	socket "github.com/AliMumtazDev/socket/models"
	socketimpl "github.com/AliMumtazDev/socket/websocket_impl"
)

var ConnMutex sync.Mutex

type WebSocketImpl struct{}

func NewWebSocketImpl() *WebSocketImpl {
	return &WebSocketImpl{}
}

func (ws *WebSocketImpl) UpgradeConnection(w http.ResponseWriter, r *http.Request) (*socketimpl.Client, error) {
	conn, err := socket.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %v", err)
		return nil, err
	}

	userID := r.URL.Query().Get("userID")

	client := &socketimpl.Client{
		Conn:   conn,
		UserID: userID,
	}

	ConnMutex.Lock()
	// socket.Connections[userID] = client
	ConnMutex.Unlock()

	return client, nil
}
