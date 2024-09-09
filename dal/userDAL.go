package dal

import (
	"th3y3m/chat-application/util"
)

func GetUserByUsername(username string) (User, error) {
	var user User

	db, err := util.ConnectToSQLServer()

	if err != nil {
		return user, err
	}

	defer db.Close()

	query := "SELECT ID, Username, Password_hash FROM Users WHERE Username = @p1"

	err = db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password_hash)

	if err != nil {
		return user, err
	}

	return user, nil
}

func GetUsers() ([]User, error) {

	db, err := util.ConnectToSQLServer()

	if err != nil {
		return nil, err
	}

	defer db.Close()

	rows, err := db.Query("SELECT ID, Username, Password_hash FROM Users")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User
	var user User

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Password_hash)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func GetUserByID(id int) (User, error) {
	var user User

	db, err := util.ConnectToSQLServer()

	if err != nil {
		return user, err
	}

	defer db.Close()

	query := "SELECT ID, Username, Password_hash FROM Users WHERE ID = @p1"

	err = db.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Password_hash)

	if err != nil {
		return user, err
	}

	return user, nil
}

func CreateUser(user User) error {
	db, err := util.ConnectToSQLServer()

	if err != nil {
		return err
	}

	defer db.Close()

	query := "INSERT INTO Users (Username, Password_hash) VALUES (@p1, @p2)"

	_, err = db.Exec(query, user.Username, user.Password_hash)

	if err != nil {
		return err
	}

	return nil
}

func UpdateUser(user User) error {
	db, err := util.ConnectToSQLServer()

	if err != nil {
		return err
	}

	defer db.Close()

	query := "UPDATE Users SET Username = @p1, Password_hash = @p2 WHERE ID = @p3"

	_, err = db.Exec(query, user.Username, user.Password_hash, user.ID)

	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(id int) error {
	db, err := util.ConnectToSQLServer()

	if err != nil {
		return err
	}

	defer db.Close()

	query := "DELETE FROM Users WHERE ID = @p1"

	_, err = db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
