package main

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

// Ports for services
const thisServerPort = "5001"
const databasePort = "5002"

// Service names for services in Kubernetes config
const databaseServiceName = "database-service"

type encryptStruct struct {
	Text string `json:"text"`
	Key  int    `json:"key"`
}

func cesar(data encryptStruct) string {

	// make all characters to lowecase
	data.Text = strings.ToLower(data.Text)
	encryptedString := ""
	// Get asciivalues for every character and add the key to shift to get the encrypted character
	for _, c := range data.Text {
		asciiValue := int(c)
		if c != ' ' {
			//fmt.Printf("Letter in for loop: %c has ASCII value: %d\n", c, asciiValue)
			shift := (asciiValue + data.Key) % 122
			if shift < 97 { // Wrap-around correction for 'z' to 'a'
				shift += 96 // Adjust the shift to wrap within lowercase letters in ASCII
			}
			encryptedChar := rune(shift)
			//fmt.Printf("Shifted character: %c\n", encryptedChar)
			encryptedString += string(encryptedChar)
		}
	}

	return encryptedString
}

// userInputHandler handles user input.
func encryptHandler(c *gin.Context) {
	fmt.Println("Handle the data to be encrypted")
	var data encryptStruct

	// Call BindJSON to bind the received JSON to encStruct.
	if err := c.BindJSON(&data); err != nil {
		fmt.Println("Error binding JSON: ", err)
		return
	}

	var encrypted encryptStruct
	encrypted.Text = cesar(data)
	encrypted.Key = data.Key

	// Respond with a JSON containing encStruct.Text and encStruct.Key
	response := gin.H{
		"text": encrypted.Text,
		"key":  encrypted.Key,
		//"error": 0,
	}

	c.JSON(200, response)
	fmt.Println("[+] Response sent to server")

}

func main() {
	router := gin.Default()

	router.POST("/encrypt", encryptHandler) // Request from server-service will be sent to this endpoint

	router.Run(":" + thisServerPort)
}
