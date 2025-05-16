package upgradeconn

import (
	"log"
	"net/http"
	"sync"

	// socket "github.com/AliMumtazDev/socket/models"
	"github.com/AliMumtazDev/socket/client"
	socket "github.com/AliMumtazDev/socket/models"
)

var ConnMutex sync.Mutex

type WebSocketImpl struct{}

func NewWebSocketImpl() *WebSocketImpl {
	return &WebSocketImpl{}
}

func (ws *WebSocketImpl) UpgradeConnection(w http.ResponseWriter, r *http.Request) (*client.Client, error) {
	conn, err := socket.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %v", err)
		return nil, err
	}

	userID := r.URL.Query().Get("userID")

	client := &client.Client{
		Conn:   conn,
		UserID: userID,
	}

	ConnMutex.Lock()
	// socket.Connections[userID] = client
	ConnMutex.Unlock()

	return client, nil
}
