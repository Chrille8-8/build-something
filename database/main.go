/*
The db looks like this
| ID (INTEGER PRIMARY KEY) | Text (TEXT)   | Key (INTEGER) |
|--------------------------|---------------|---------------|
| 1                        | afnasncasn    | 3             |
| 2                        | xjhhnedkjv    | 55            |
|--------------------------|---------------|---------------|
*/

package main

import (
	//"time"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

const thisServerPort = "5002"

type data struct {
	ID   int
	Text string
	Key  int
}

// Gets everything in the database with the id
func get(c *gin.Context) {

	id := c.Param("id")
	fmt.Printf("id: %s\n", id)
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Select everything from text and
	cipher_from_db, err := db.Query("SELECT * FROM chiphers WHERE ID = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer cipher_from_db.Close()
	fmt.Println("cipher_from_db: ", cipher_from_db)
	var item data
	err = db.QueryRow("SELECT ID, Text, Key FROM chiphers WHERE ID = ?", id).Scan(&item.ID, &item.Text, &item.Key)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
			return
		}
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, item)
}

func getAll(c *gin.Context) {

	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Select everything from text and
	res_from_db, err := db.Query("SELECT * FROM chiphers")
	if err != nil {
		log.Fatal(err)
	}
	defer res_from_db.Close()
	var results []data // Slice to store fetched data

	for res_from_db.Next() { // Iterate and fetch the records from result cursor
		item := data{}
		err := res_from_db.Scan(&item.ID, &item.Text, &item.Key)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, item) // Append each row to the results slice
	}
	db.Close()

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, results)
}

// Insert to database
func putDatabase(c *gin.Context) {

	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var d data

	// Call BindJSON to bind the received JSON to encStruct.
	if err := c.BindJSON(&d); err != nil {
		fmt.Println("Error binding JSON:", err)
		return
	}

	_, err = db.Exec("INSERT INTO chiphers (Text, Key) VALUES (?, ?)", d.Text, d.Key)
	if err != nil {
		log.Fatal(err)
	}

	db.Close()
	c.JSON(http.StatusOK, "PUT successful!")
}

func removeDatabase(c *gin.Context) {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, "Something went wrong!")

	}
	defer db.Close()

	id := c.Param("id")
	_, err = db.Exec("DELETE FROM chiphers WHERE ID = ?", id)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, "Something went wrong!")
	}

	db.Close()
	c.JSON(http.StatusOK, "REMOVE successful!")
}

// Insert test data into the database
func addTest(c *gin.Context) {
	testData := []struct {
		Text string
		Key  int
	}{
		{"test1", 123},
		{"test2", 456},
		{"test3", 789},
	}

	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO chiphers (Text, Key) VALUES (?, ?)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()

	for _, data := range testData {
		_, err := stmt.Exec(data.Text, data.Key)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, "Test data inserted successfully!")
}

func main() {

	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		fmt.Println("Could not open db")
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS chiphers (ID INTEGER PRIMARY KEY, Text TEXT, Key INTEGER)")
	if err != nil {
		fmt.Println("Could not Create table")
		log.Fatal(err)
	}
	db.Close()

	router := gin.Default()

	// Route to get the chipher of id
	router.GET("/get/:id", get)

	// Gets all rows in the database
	router.GET("/getall", getAll) // This can be called from the server-service

	// Route to add the encrypted text and the key used for encryption
	router.POST("/add", putDatabase) // This can be called from the server-service

	// Romeves the row with the specified id
	router.GET("/remove/:id", removeDatabase)

	// Adds test data to DB
	router.GET("/add_test", addTest)

	// Run the server
	router.Run(":" + thisServerPort)
}
