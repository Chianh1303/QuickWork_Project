package handlers

import (
	"QuickWork/models"
	"encoding/json"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
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

			Clients: make(map[uint]*websocket.Conn),

			Register: make(chan *ClientBind),

			Unregister: make(chan uint),

			DB: db,
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

		// Gán người gửi
		message.SenderID = userID
		message.CreatedAt = time.Now()

		// Lưu DB
		if err := GlobalChatHub.DB.Create(&message).Error; err != nil {
			log.Println("Save Error:", err)
			continue
		}

		// Sau khi lưu thì ID sẽ được GORM sinh
		messageJSON, _ := json.Marshal(message)

		// Gửi cho người gửi và người nhận
		GlobalChatHub.Mu.RLock()

		// Sender
		if senderConn, ok := GlobalChatHub.Clients[message.SenderID]; ok {
			log.Println("SEND TO SELF")
			senderConn.WriteMessage(websocket.TextMessage, messageJSON)
		}

		// Receiver
		if receiverConn, ok := GlobalChatHub.Clients[message.ReceiverID]; ok {
			log.Println("SEND TO RECEIVER")
			receiverConn.WriteMessage(websocket.TextMessage, messageJSON)
		}

		GlobalChatHub.Mu.RUnlock()
	}
}

func GetChatHistory(db *gorm.DB) fiber.Handler {

	return func(c *fiber.Ctx) error {

		appIDStr := c.Query("application_id")

		appID, err := strconv.Atoi(appIDStr)

		if err != nil {

			return c.Status(400).JSON(fiber.Map{

				"message": "application_id invalid",
			})

		}

		var messages []models.Message

		if err := db.

			Where("application_id = ?", appID).

			Order("created_at asc").

			Find(&messages).Error; err != nil {

			return c.Status(500).JSON(fiber.Map{

				"message": err.Error(),
			})

		}

		return c.JSON(messages)

	}

}