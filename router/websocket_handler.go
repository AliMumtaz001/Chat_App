package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AliMumtazDev/Go_Chat_App/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var userMsg models.Message

func (r *Router) WebSocketHandler(c *gin.Context) {
	userID := c.Query("userID")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing userID"})
		return
	}

	mutex.Lock()
	if _, exists := connections[userID]; exists {
		mutex.Unlock()
		fmt.Println(http.StatusConflict, gin.H{"error": "Connection already exists for this user"})
		return
	}

	mutex.Unlock()
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upgrade connection"})
		return
	}

	mutex.Lock()
	connections[userID] = conn
	mutex.Unlock()
	fmt.Printf("User %s connected\n", userID)

	defer func() {
		mutex.Lock()
		delete(connections, userID)
		mutex.Unlock()
		conn.Close()
		fmt.Printf("User %s disconnected\n", userID)
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		var message models.WebSocketMessage
		if err := json.Unmarshal(msg, &message); err != nil {
			conn.WriteMessage(websocket.TextMessage, []byte("Invalid message format"))
			continue
		}
		fmt.Printf("Received message: %s\n", msg)

		switch message.Type {
		case "sendmessage":
			var dbMessage models.Message
			dbMessage.SenderID, _ = strconv.ParseInt(userID, 10, 64)
			dbMessage.ReceiverID, err = strconv.ParseInt(message.To, 10, 64)
			if err != nil {
				conn.WriteMessage(websocket.TextMessage, []byte("Invalid recipient ID format"))
				continue
			}
			dbMessage.Content = message.Content

			err := r.UserService.SendMessageservice(c, dbMessage.SenderID, dbMessage.ReceiverID, dbMessage)
			if err != nil {
				conn.WriteMessage(websocket.TextMessage, []byte("Failed to save message"))
				continue
			}

			mutex.Lock()
			recipientConn, exists := connections[message.To]
			mutex.Unlock()
			if exists {
				err := recipientConn.WriteMessage(websocket.TextMessage, msg)
				if err != nil {
					conn.WriteMessage(websocket.TextMessage, []byte("Failed to deliver message"))
				}
			}
			fmt.Printf("Message sent to recipient %s\n", message.To)

		case "getmessage":
			senderID := userID
			receiverID := message.To
			messages, err := r.UserService.GetMessageservice(c, senderID, receiverID)
			if err != nil {
				conn.WriteMessage(websocket.TextMessage, []byte("Failed to fetch messages: "+err.Error()))
				continue
			}
			messagesJSON, err := json.Marshal(messages)
			if err != nil {
				conn.WriteMessage(websocket.TextMessage, []byte("Failed to encode messages: "+err.Error()))
				continue
			}
			conn.WriteMessage(websocket.TextMessage, messagesJSON)
		default:
			conn.WriteMessage(websocket.TextMessage, []byte("Unknown message type"))
		}
	}
}
