// main.go
package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// Show the index.html
func showInputForm(c *gin.Context) {
	filePath := filepath.Join("html", "index.html")
	c.File(filePath)
}

// Takes the input from user and sends it to encryption server and then displays it
func processData(c *gin.Context) {
	text := c.PostForm("text")
	number := c.PostForm("number")

	fmt.Printf("Text: %s, Number: %s\n", text, number)

	//
	c.String(http.StatusOK, "Received data. Check server terminal for details.")
}

func main() {
	router := gin.Default()

	// Serve static files from the "html" directory
	router.Static("/static", "./html")

	// Route to where text and key will be entered
	router.GET("/input", showInputForm)

	// Route to display the
	router.POST("/submit", processData)

	// Run the server
	router.Run(":8080")
}
