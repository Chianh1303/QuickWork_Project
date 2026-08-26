package websocket

import (
	"encoding/json"
	"log"
	"strconv"
	"sync"
	"time"

	"QuickWork/internal/models"

	"github.com/gofiber/websocket/v2"
	"gorm.io/gorm"
)

type ClientBind struct {
	UserID uint
	Conn   *websocket.Conn
}

type ChatHub struct {
	Clients    map[uint]*websocket.Conn
	Register   chan *ClientBind
	Unregister chan uint
	Mu         sync.RWMutex
	DB         *gorm.DB
}

var (
	GlobalChatHub *ChatHub
	once          sync.Once
)

func StartChatHub(db *gorm.DB) *ChatHub {
	once.Do(func() {
		GlobalChatHub = &ChatHub{
			Clients:    make(map[uint]*websocket.Conn),
			Register:   make(chan *ClientBind),
			Unregister: make(chan uint),
			DB:         db,
		}
		go GlobalChatHub.run()
	})
	return GlobalChatHub
}

func (h *ChatHub) run() {
	for {
		select {
		case client := <-h.Register:
			h.Mu.Lock()
			h.Clients[client.UserID] = client.Conn
			h.Mu.Unlock()
			log.Printf("User %d connected\n", client.UserID)
		case userID := <-h.Unregister:
			h.Mu.Lock()
			if conn, ok := h.Clients[userID]; ok {
				conn.Close()
				delete(h.Clients, userID)
				log.Printf("User %d disconnected\n", userID)
			}
			h.Mu.Unlock()
		}
	}
}

func HandleWS(conn *websocket.Conn) {
	userIDStr := conn.Query("userId")
	userID64, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		conn.Close()
		return
	}

	userID := uint(userID64)
	GlobalChatHub.Register <- &ClientBind{
		UserID: userID,
		Conn:   conn,
	}

	defer func() {
		GlobalChatHub.Unregister <- userID
	}()

	for {
		_, messageBytes, err := conn.ReadMessage()
		if err != nil {
			break
		}

		var message models.Message
		if err := json.Unmarshal(messageBytes, &message); err != nil {
			log.Println("JSON Error:", err)
			continue
		}

		message.SenderID = userID
		message.CreatedAt = time.Now()

		if err := GlobalChatHub.DB.Create(&message).Error; err != nil {
			log.Println("Save Error:", err)
			continue
		}

		// Tạo Notification cho người nhận khi có tin nhắn mới
		if message.ReceiverID > 0 {
			senderName := "Người dùng"
			var student models.Student
			if err := GlobalChatHub.DB.Where("user_id = ?", message.SenderID).First(&student).Error; err == nil && student.FullName != "" {
				senderName = student.FullName
			} else {
				var business models.Business
				if err := GlobalChatHub.DB.Where("user_id = ?", message.SenderID).First(&business).Error; err == nil && business.CompanyName != "" {
					senderName = business.CompanyName
				}
			}

			notifMessage := message.MessageText
			if len(notifMessage) > 100 {
				notifMessage = notifMessage[:97] + "..."
			}

			notif := models.Notification{
				UserID:      message.ReceiverID,
				Title:       "💬 Tin nhắn mới từ " + senderName,
				Message:     notifMessage,
				Type:        "chat",
				ReferenceID: message.ApplicationID,
				IsRead:      false,
				CreatedAt:   time.Now(),
			}

			if err := GlobalChatHub.DB.Create(&notif).Error; err != nil {
				log.Println("Failed to create chat notification:", err)
			} else {
				log.Printf("🔔 [Chat Notification Created]: Sent to Receiver UserID=%d from %s (AppID=%d)", message.ReceiverID, senderName, message.ApplicationID)
			}
		}

		messageJSON, _ := json.Marshal(message)

		GlobalChatHub.Mu.RLock()
		if senderConn, ok := GlobalChatHub.Clients[message.SenderID]; ok {
			log.Println("SEND TO SELF")
			senderConn.WriteMessage(websocket.TextMessage, messageJSON)
		}
		if receiverConn, ok := GlobalChatHub.Clients[message.ReceiverID]; ok {
			log.Println("SEND TO RECEIVER")
			receiverConn.WriteMessage(websocket.TextMessage, messageJSON)
		}
		GlobalChatHub.Mu.RUnlock()
	}
}
