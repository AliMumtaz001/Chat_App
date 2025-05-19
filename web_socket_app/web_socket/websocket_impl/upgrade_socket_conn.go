package websocket_impl

import (
	"log"
	"strconv"
	"sync"

	userserviceimpl "github.com/AliMumtazDev/Go_Chat_App/api/message_service"
	"github.com/AliMumtazDev/socket/client"
	"github.com/gin-gonic/gin"
)

var ConnLock = sync.Mutex{}
type WebSocketImpl struct{}
func NewWebSocketImpl() *WebSocketImpl {
	return &WebSocketImpl{}
}

func (w *WebSocketServiceImpl) AddConn(userID string, wsConn *client.Client, c *gin.Context) error {
	uid, _ := strconv.Atoi(userID)
	log.Println("User ID:", uid)
	ConnLock.Lock()
	w.Clients[uid] = wsConn
	ConnLock.Unlock()

	log.Println("Conn", w.Clients)

	log.Println("User connected:", uid)

	defer func() {
		ConnLock.Lock()
		delete(w.Clients, uid)
		ConnLock.Unlock()
		wsConn.Conn.Close()
		log.Println("User disconnected:", uid)
	}()

	for {
		var incoming userserviceimpl.ServerMesageToSocket
		err := wsConn.Conn.ReadJSON(&incoming)
		if err != nil {
			log.Println("Error reading JSON:", err)
			break
		}

		log.Println("Received JSON from", uid, incoming)

		action := incoming.Action

		if action == "send" {
			log.Println("Sending message to", uid)
			receiverIDFloat := incoming.DestinationID
			receiverID := int(receiverIDFloat)

			message := incoming.Content
			if conn, ok := w.Clients[receiverID]; ok {

				err := conn.Conn.WriteJSON(map[string]any{
					"receiverID": uid,
					"message":    message,
				})

				if err != nil {
					log.Println("Error writing JSON to receiver:", err)
				}
			} else {
				log.Println("Receiver not connected:", receiverID)
			}
		}
	}

	return nil
}