package dal

import (
	"th3y3m/chat-application/util"
)

// GetRooms retrieves all rooms from the database
func GetRooms() ([]Room, error) {

	db, err := util.ConnectToSQLServer()

	if err != nil {
		return nil, err
	}

	defer db.Close()

	rows, err := db.Query("SELECT RoomId, RoomName FROM Rooms")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var rooms []Room
	var room Room

	for rows.Next() {
		err := rows.Scan(&room.RoomId, &room.RoomName)

		if err != nil {
			return nil, err
		}

		rooms = append(rooms, room)
	}

	return rooms, nil
}

// GetRoomByID retrieves a room by its ID from the database
func GetRoomByID(id int) (Room, error) {
	var room Room

	db, err := util.ConnectToSQLServer()

	if err != nil {
		return room, err
	}

	defer db.Close()

	query := "SELECT RoomId, RoomName FROM Rooms WHERE RoomId = @p1"

	err = db.QueryRow(query, id).Scan(&room.RoomId, &room.RoomName)

	if err != nil {
		return room, err
	}

	return room, nil
}

// CreateRoom creates a new room in the database
func CreateRoom(room Room) error {
	db, err := util.ConnectToSQLServer()

	if err != nil {
		return err
	}

	defer db.Close()

	query := "INSERT INTO Rooms (RoomName) VALUES (@p1)"

	_, err = db.Exec(query, room.RoomName)

	if err != nil {
		return err
	}

	return nil
}

// UpdateRoom updates an existing room in the database
func UpdateRoom(room Room) error {
	db, err := util.ConnectToSQLServer()

	if err != nil {
		return err
	}

	defer db.Close()

	query := "UPDATE Rooms SET RoomName = @p1 WHERE RoomId = @p2"

	_, err = db.Exec(query, room.RoomName, room.RoomId)

	if err != nil {
		return err
	}

	return nil
}

// DeleteRoom deletes a room by its ID from the database
func DeleteRoom(id int) error {
	db, err := util.ConnectToSQLServer()

	if err != nil {
		return err
	}

	defer db.Close()

	query := "DELETE FROM Rooms WHERE RoomId = @p1"

	_, err = db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
