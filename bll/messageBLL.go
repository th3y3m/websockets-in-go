package bll

import (
	"th3y3m/chat-application/dal"
	"time"
)

// GetMessages returns all messages
func GetMessages() ([]dal.Message, error) {
	return dal.GetMessages()
}

// GetMessageByID returns a message by its ID
func GetMessageByID(id int) (dal.Message, error) {
	return dal.GetMessageByID(id)
}

// CreateMessage creates a new message
func CreateMessage(newMessage dal.NewMessage) error {
	message := dal.Message{
		RoomId:    newMessage.RoomId,
		UserId:    newMessage.UserId,
		Message:   newMessage.Message,
		CreatedAt: time.Now(),
	}
	return dal.CreateMessage(message)
}

// UpdateMessage updates a message
func UpdateMessage(id int, messageContent string) error {
	message, err := GetMessageByID(id)
	if err != nil {
		return err
	}

	message.Message = messageContent

	return dal.UpdateMessage(message)
}

// DeleteMessage deletes a message
func DeleteMessage(id int) error {
	return dal.DeleteMessage(id)
}

// GetMessagesByRoomID returns all messages in a room
func GetMessagesByRoomID(roomID int) ([]dal.Message, error) {
	return dal.GetMessageByRoomID(roomID)
}
