package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"github.com/AliMumtazDev/Go_Chat_App/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var connections = make(map[string]*websocket.Conn)
var mutex = &sync.Mutex{}

func (r *Router) WebSocketHandler(c *gin.Context) {
	userID := c.Query("userID")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing userID"})
		return
	}

	mutex.Lock()
	if _, exists := connections[userID]; exists {
		mutex.Unlock()
		c.JSON(http.StatusConflict, gin.H{"error": "Connection already exists for this user"})
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

	defer func() {
		mutex.Lock()
		delete(connections, userID)
		mutex.Unlock()
		conn.Close()
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

		message.From = userID

		switch message.Type {
		case "sendmessage":
			sID, err := strconv.ParseInt(userID, 10, 64)
			if err != nil {
				conn.WriteMessage(websocket.TextMessage, []byte("Invalid userID format"))
				continue
			}
			rID, err := strconv.ParseInt(message.To, 10, 64)
			if err != nil {
				conn.WriteMessage(websocket.TextMessage, []byte("Invalid recipient ID format"))
				continue
			}
			userID, err := strconv.ParseInt(userID, 10, 64)
			if err != nil {
				conn.WriteMessage(websocket.TextMessage, []byte("Invalid userID format"))
			}
			
			dbMessage := models.Message{
				SenderID:   userID,
				ReceiverID: rID,
				Content:    message.Content,
			}
			err = r.UserService.SendMessageservice(c, sID, rID, dbMessage)
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

			conn.WriteMessage(websocket.TextMessage, msg)

		case "getmessage":
			messages, err := r.UserService.GetMessageservice(c, userID, message.To)
			if err != nil {
				conn.WriteMessage(websocket.TextMessage, []byte("Failed to fetch messages"))
				continue
			}

			messagesJSON, err := json.Marshal(messages)
			if err != nil {
				conn.WriteMessage(websocket.TextMessage, []byte("Failed to encode messages"))
				continue
			}
			conn.WriteMessage(websocket.TextMessage, messagesJSON)

		default:
			conn.WriteMessage(websocket.TextMessage, []byte("Unknown message type"))
		}
	}
}
