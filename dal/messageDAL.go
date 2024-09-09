package dal

import (
	"th3y3m/chat-application/util"
)

func GetMessages() ([]Message, error) {

	db, err := util.ConnectToSQLServer()

	if err != nil {
		return nil, err
	}

	defer db.Close()

	rows, err := db.Query("SELECT MessageId, RoomId, UserId, Message FROM Messages")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var messages []Message
	var message Message

	for rows.Next() {
		err := rows.Scan(&message.MessageId, &message.RoomId, &message.UserId, &message.Message)

		if err != nil {
			return nil, err
		}

		messages = append(messages, message)
	}

	return messages, nil
}

func GetMessageByID(id int) (Message, error) {
	var message Message

	db, err := util.ConnectToSQLServer()

	if err != nil {
		return message, err
	}

	defer db.Close()

	query := "SELECT MessageId, RoomId, UserId, Message FROM Messages WHERE MessageId = @p1"

	err = db.QueryRow(query, id).Scan(&message.MessageId, &message.RoomId, &message.UserId, &message.Message)

	if err != nil {
		return message, err
	}

	return message, nil
}

func CreateMessage(message Message) error {
	db, err := util.ConnectToSQLServer()

	if err != nil {
		return err
	}

	defer db.Close()

	query := "INSERT INTO Messages (RoomId, UserId, Message) VALUES (@p1, @p2, @p3)"

	_, err = db.Exec(query, message.RoomId, message.UserId, message.Message)

	if err != nil {
		return err
	}

	return nil
}

func UpdateMessage(message Message) error {
	db, err := util.ConnectToSQLServer()

	if err != nil {
		return err
	}

	defer db.Close()

	query := "UPDATE Messages SET RoomId = @p1, UserId = @p2, Message = @p3 WHERE MessageId = @p4"

	_, err = db.Exec(query, message.RoomId, message.UserId, message.Message, message.MessageId)

	if err != nil {
		return err
	}

	return nil
}

func DeleteMessage(id int) error {
	db, err := util.ConnectToSQLServer()

	if err != nil {
		return err
	}

	defer db.Close()

	query := "DELETE FROM Messages WHERE MessageId = @p1"

	_, err = db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

func GetMessageByRoomID(id int) ([]Message, error) {
	db, err := util.ConnectToSQLServer()

	if err != nil {
		return nil, err
	}

	defer db.Close()

	query := "SELECT MessageId, RoomId, UserId, Message FROM Messages WHERE RoomId = @p1"

	rows, err := db.Query(query, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var messages []Message
	var message Message

	for rows.Next() {
		err := rows.Scan(&message.MessageId, &message.RoomId, &message.UserId, &message.Message)

		if err != nil {
			return nil, err
		}

		messages = append(messages, message)
	}

	return messages, nil
}
