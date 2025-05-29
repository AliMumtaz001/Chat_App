package websocket_impl

import (
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/AliMumtazDev/Go_Chat_App/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var ConnLock = sync.Mutex{}

func (w *WebSocketServiceImpl) AddConn(userID string, wsConn *websocket.Conn, c *gin.Context) error {
	uid, err := strconv.Atoi(userID)
	if err != nil {
		log.Println("Error converting userID to int:", err)
		return err
	}
	log.Println("User ID:", uid)
	ConnLock.Lock()
	w.Clients[uid] = wsConn
	ConnLock.Unlock()
	fmt.Println("connected clients:", w.Clients)
	fmt.Println("connected clientshghg:", wsConn)
	log.Println("User connected:", uid)
	go func() {
		defer func() {
			ConnLock.Lock()
			delete(w.Clients, uid)
			ConnLock.Unlock()
			wsConn.Close()
			log.Println("User disconnected:", uid)
		}()
		for {
			// var incoming userserviceimpl.models.ServerMesageToSocket
			var incoming models.ServerMesageToSocket
			err := wsConn.ReadJSON(&incoming)
			if err != nil {
				log.Println("Error reading JSON:", err)
				break
			}
			log.Println("Received JSON from", uid, ":", incoming)
			if incoming.Action == "send" {
				receiverID := int(incoming.DestinationID)
				log.Printf("Received message for receiver_id: %d", incoming.DestinationID)
				if receiverID == uid {
					log.Println("Cannot send message to self")
					continue
				}
				ConnLock.Lock()
				conn, ok := w.Clients[receiverID]
				ConnLock.Unlock()
				if ok {
					message := map[string]interface{}{
						"sender_id": uid,
						"content":   incoming.Content,
					}
					err := conn.WriteJSON(message)
					if err != nil {
						log.Println("Error writing JSON to receiver:", err)
					} else {
						log.Println("Message sent to receiver:", receiverID)
					}
				} else {
					log.Println("Receiver not connected:", receiverID)
				}
			}
		}
	}()
	return nil
}
