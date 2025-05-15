// package socket

// import (
// 	"log"
// 	"net/http"
// )

// func (impl WebSocketImpl) RegisterWebSocketRoute(mux *http.ServeMux, w http.ResponseWriter, r *http.Request) {
// 	ws := NewWebSocketImpl()
// 	client, err := ws.UpgradeConnection(w, r)
// 	if err != nil {
// 		http.Error(w, "Failed to upgrade connection", http.StatusInternalServerError)
// 		return
// 	}
// 	for {
// 		message, err := ws.ReceiveMessage(client)
// 		if err != nil {
// 			log.Printf("Error reading message: %v", err)
// 			return
// 		}
// 		// Echo the message back to the client (for simplicity)
// 		err = ws.SendMessage(client, message)
// 		if err != nil {
// 			log.Printf("Error sending message: %v", err)
// 			return
// 		}
// 	}
// }

package socketimpl

import (
	"log"
	"net/http"

	upgradeconn "github.com/AliMumtazDev/Go_Chat_App/web_socket_app/database"
	socketimpl "github.com/AliMumtazDev/socket/websocket_impl"
	"github.com/gin-gonic/gin"
)

func (impl *upgradeconn.WebSocketImpl) RegisterWebSocketRoute(c *gin.Context) {
	ws := socketimpl.NewWebSocketImpl()
	//NewWebSocketImpl()
	client, err := ws.UpgradeConnection(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upgrade connection"})
		return
	}
	for {
		message, err := ws.ReceiveMessage(client)
		if err != nil {
			log.Printf("Error reading message: %v", err)
			return
		}
		err = ws.SendMessage(client, message)
		if err != nil {
			log.Printf("Error sending message: %v", err)
			return
		}
	}
}
