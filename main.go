package main

import (
	"go-lang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.GET("/", models.SendSomeData)
	router.GET("/get-data", getData)
	router.GET("/get-data/:id", getDataById)
	router.POST("/create-data", createData)
	router.Run("localhost:8085")
}

func getData(c *gin.Context) {
	data := models.GetAllData()
	if data == nil || len(data) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "Nothing found",
		})
	} else {
		c.JSON(http.StatusOK, data)
	}
}

func getDataById(c *gin.Context) {
	id := c.Param("id")
	data := models.GetDataById(id)

	if data == nil || len(data) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "Nothing found",
		})
	} else {
		c.JSON(http.StatusOK, data)
	}
}

func createData(c *gin.Context) {
	var data models.Data

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Somthing went wrong",
		})
	} else {
		models.CreateData(data)
		c.JSON(http.StatusCreated, data)
	}
}
