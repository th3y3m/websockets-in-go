package bll

import (
	"th3y3m/chat-application/dal"
)

// GetRooms returns all rooms
func GetRooms() ([]dal.Room, error) {
	return dal.GetRooms()
}

// GetRoomByID returns a room by its ID
func GetRoomByID(id int) (dal.Room, error) {
	return dal.GetRoomByID(id)
}

// CreateRoom creates a new room
func CreateRoom(roomName string) error {
	room := dal.Room{
		RoomName: roomName,
	}
	return dal.CreateRoom(room)
}

// UpdateRoom updates a room
func UpdateRoom(roomName string) error {
	room := dal.Room{
		RoomName: roomName,
	}
	return dal.UpdateRoom(room)
}

// DeleteRoom deletes a room
func DeleteRoom(id int) error {
	return dal.DeleteRoom(id)
}
