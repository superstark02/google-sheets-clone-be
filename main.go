package main

import (
	"go-lang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.Use(CORSMiddleware())

	router.GET("/", models.SendSomeData)
	router.GET("/get-data", getData)
	router.GET("/get-data/:id", getDataById)
	router.POST("/create-data", createData)
	router.Run("localhost:8085")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
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
