package models

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

const dbLocation = "root:R3tain$55@tcp(127.0.0.1:3306)/GOOGLE_SHEETS"

func GetAllData() []Data {
	db, err := sql.Open("mysql", dbLocation)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	defer db.Close()

	//select, err := db.Query("use GOOGLE_SHEETS;")
	results, err := db.Query("SELECT * FROM sheet_data")

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	response := []Data{}
	for results.Next() {
		var data Data
		err = results.Scan(&data.Id, &data.Color, &data.Data, &data.Bold, &data.Y, &data.X)

		if err != nil {
			panic(err.Error())
		}
		response = append(response, data)
	}
	return response
}

func GetDataByRange() []Data {
	db, err := sql.Open("mysql", dbLocation)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM sheet_data where id = ")

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	response := []Data{}
	for results.Next() {
		var data Data
		err = results.Scan(&data.Id, &data.Bold, &data.Data, &data.Color)

		if err != nil {
			panic(err.Error())
		}
		response = append(response, data)
	}
	return response
}

func GetDataById(id string) []Data {
	db, err := sql.Open("mysql", dbLocation)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM sheet_data where id = ?", id)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	response := []Data{}
	for results.Next() {
		var data Data
		err = results.Scan(&data.Id, &data.Bold, &data.Data, &data.Color)

		if err != nil {
			panic(err.Error())
		}
		response = append(response, data)
	}
	return response
}

func SendSomeData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from gin",
	})
}

func CreateData(data Data) {
	db, err := sql.Open("mysql", dbLocation)
	if err != nil {
		fmt.Println("Err", err.Error())
	}

	defer db.Close()
	insert, err := db.Query(
		"INSERT INTO sheet_data (id,data,bold,color) VALUES (?,?,?,?)",
		data.Id, data.Data, data.Bold, data.Color)

	if err != nil {
		fmt.Println("Err", err.Error())
	}
	defer insert.Close()
}
