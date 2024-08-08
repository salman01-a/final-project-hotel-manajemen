package controllers

import (
	"final-project-sanbercode/database"
	"final-project-sanbercode/models"
	"final-project-sanbercode/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllPayment(c *gin.Context) {
	var (
		result gin.H
	)

	payment, err := repository.GetAllPayment(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": payment,
		}
	}
	c.JSON(http.StatusOK, result)
}

func InsertPayment(c *gin.Context) {
	var payment models.Payment
	err := c.BindJSON(&payment)
	if err != nil {
		panic(err)
	}
	err = repository.InsertPayment(database.DbConnection, payment)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, payment)
}
func UpdatePayment(c *gin.Context) {
	var payment models.Payment
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&payment)
	if err != nil {
		panic(err)
	}
	payment.ID = id
	err = repository.UpdatePayment(database.DbConnection, payment)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, payment)
}
func DeletePayment(c *gin.Context) {
	var payment models.Payment
	id, _ := strconv.Atoi(c.Param("id"))
	payment.ID = id
	err := repository.DeletePayment(database.DbConnection, payment)

	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, payment)
}
