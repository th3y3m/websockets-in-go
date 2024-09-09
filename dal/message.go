package dal

import "time"

// Message represents a chat message in the database
type Message struct {
	MessageId int       `json:"message_id"`
	RoomId    int       `json:"room_id"`
	UserId    int       `json:"user_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}
type NewMessage struct {
	RoomId  int    `json:"room_id"`
	UserId  int    `json:"user_id"`
	Message string `json:"message"`
}
