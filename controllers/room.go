package controllers

import (
	"final-project-sanbercode/database"
	"final-project-sanbercode/models"
	"final-project-sanbercode/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllRoom(c *gin.Context) {
	var (
		result gin.H
	)

	room, err := repository.GetAllRoom(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": room,
		}
	}
	c.JSON(http.StatusOK, result)
}

func InsertRoom(c *gin.Context) {
	var room models.Room
	err := c.BindJSON(&room)
	if err != nil {
		panic(err)
	}
	err = repository.InsertRoom(database.DbConnection, room)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, room)
}
func UpdateRoom(c *gin.Context) {
	var room models.Room
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&room)
	if err != nil {
		panic(err)
	}
	room.ID = id
	err = repository.UpdateRoom(database.DbConnection, room)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, room)
}
func DeleteRoom(c *gin.Context) {
	var room models.Room
	id, _ := strconv.Atoi(c.Param("id"))
	room.ID = id
	err := repository.DeleteRoom(database.DbConnection, room)

	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, room)
}
func GetRoomByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	room, err := repository.GetRoomByID(database.DbConnection, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"result": "Room not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": room})
}
