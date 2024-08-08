package controllers

import (
	"final-project-sanbercode/database"
	"final-project-sanbercode/models"
	"final-project-sanbercode/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllBooking(c *gin.Context) {
	var (
		result gin.H
	)

	booking, err := repository.GetAllBooking(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": booking,
		}
	}
	c.JSON(http.StatusOK, result)
}

func InsertBooking(c *gin.Context) {
	var booking models.Booking
	err := c.BindJSON(&booking)
	if err != nil {
		panic(err)
	}
	err = repository.InsertBooking(database.DbConnection, booking)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, booking)
}
func UpdateBooking(c *gin.Context) {
	var booking models.Booking
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&booking)
	if err != nil {
		panic(err)
	}
	booking.ID = id
	err = repository.UpdateBooking(database.DbConnection, booking)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, booking)
}
func DeleteBooking(c *gin.Context) {
	var booking models.Booking
	id, _ := strconv.Atoi(c.Param("id"))
	booking.ID = id
	err := repository.DeleteBooking(database.DbConnection, booking)

	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, booking)
}
