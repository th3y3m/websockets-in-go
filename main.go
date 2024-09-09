package main

import (
	"fmt"
	"log"
	"net/http"
	"th3y3m/chat-application/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type RoomHost struct {
	clients   map[*websocket.Conn]bool
	broadcast chan []byte
}

var rooms = make(map[string]*RoomHost)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins
	},
}

func main() {
	router := gin.Default()

	// Configure CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Adjust this to your frontend's origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// WebSocket route
	router.GET("/ws", func(c *gin.Context) {
		handleConnections(c.Writer, c.Request)
	})

	// User routes
	router.POST("/login", api.Login)
	router.POST("/register", api.Register)
	router.GET("/users", api.GetUsers)
	router.GET("/users/:id", api.GetUserByID)
	router.PUT("/users/:id", api.UpdateUser)
	router.DELETE("/users/:id", api.DeleteUser)

	// Message routes
	router.GET("/messages", api.GetMessages)
	router.GET("/messages/:id", api.GetMessageByID)
	router.POST("/messages", api.CreateMessage)
	router.PUT("/messages/:id", api.UpdateMessage)
	router.DELETE("/messages/:id", api.DeleteMessage)
	router.GET("/messages/room/:id", api.GetMessagesByRoomID)

	// Room routes
	router.GET("/rooms", api.GetRooms)
	router.GET("/rooms/:id", api.GetRoomByID)
	router.POST("/rooms", api.CreateRoom)
	router.PUT("/rooms/:id", api.UpdateRoom)
	router.DELETE("/rooms/:id", api.DeleteRoom)

	// Start the server on port 8080
	fmt.Println("Server started at :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Get the room key from the URL query parameters
	roomKey := r.URL.Query().Get("room")
	if roomKey == "" {
		http.Error(w, "Room key is required", http.StatusBadRequest)
		return
	}

	// Upgrade initial GET request to a WebSocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer ws.Close()

	// Ensure the room exists
	if _, ok := rooms[roomKey]; !ok {
		rooms[roomKey] = &RoomHost{
			clients:   make(map[*websocket.Conn]bool),
			broadcast: make(chan []byte),
		}
		go handleMessages(roomKey)
	}

	// Register the new client in the room
	room := rooms[roomKey]
	room.clients[ws] = true

	// Listen for incoming messages
	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			delete(room.clients, ws)
			break
		}
		// Send the message to the broadcast channel
		room.broadcast <- msg
	}
}

func handleMessages(roomKey string) {
	room := rooms[roomKey]
	for {
		msg := <-room.broadcast
		// Send the message to all clients in the room
		for client := range room.clients {
			err := client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println("Error writing message:", err)
				client.Close()
				delete(room.clients, client)
			}
		}
	}
}
