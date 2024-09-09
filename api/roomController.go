package api

import (
	"net/http"
	"strconv"
	"th3y3m/chat-application/bll"

	"github.com/gin-gonic/gin"
)

func GetRooms(c *gin.Context) {
	rooms, err := bll.GetRooms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"rooms": rooms})
}

func GetRoomByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Invalid": err.Error()})
		return
	}
	room, err := bll.GetRoomByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"room": room})
}

func CreateRoom(c *gin.Context) {

	roomName := c.PostForm("roomName")
	err := bll.CreateRoom(roomName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Room created"})
}

func UpdateRoom(c *gin.Context) {
	roomName := c.PostForm("roomName")
	err := bll.UpdateRoom(roomName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Room updated"})
}

func DeleteRoom(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Invalid": err.Error()})
		return
	}
	err1 := bll.DeleteRoom(id)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Room deleted"})
}
