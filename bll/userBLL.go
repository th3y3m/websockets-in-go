package bll

import (
	"th3y3m/chat-application/dal"
	"th3y3m/chat-application/util"

	"golang.org/x/crypto/bcrypt"
)

func Login(username, password string) (string, error) {
	user, err := dal.GetUserByUsername(username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password_hash), []byte(password))
	if err != nil {
		return "", err
	}
	token, err := util.GenerateToken(user.Username)

	if err != nil {
		return "", err
	}

	return token, nil
}

func GetUsers() ([]dal.User, error) {
	return dal.GetUsers()
}

func GetUserByID(id int) (dal.User, error) {
	return dal.GetUserByID(id)
}

func UpdateUser(id int, username, password string) error {
	updateUser, err := GetUserByID(id)
	if err != nil {
		return err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	updateUser.Username = username
	updateUser.Password_hash = string(hash)

	return dal.UpdateUser(updateUser)
}

func DeleteUser(id int) error {
	return dal.DeleteUser(id)
}

func GetUserByUsername(username string) (dal.User, error) {
	return dal.GetUserByUsername(username)
}

func CreateUser(user dal.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password_hash), 14)
	if err != nil {
		return err
	}
	user.Password_hash = string(hash)
	return dal.CreateUser(user)
}
