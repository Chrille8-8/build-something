package main

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

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
			fmt.Printf("Letter in for loop: %c has ASCII value: %d\n", c, asciiValue)
			shift := (asciiValue + data.Key) % 122
			if shift < 97 { // Wrap-around correction for 'z' to 'a'
				shift += 96 // Adjust the shift to wrap within lowercase letters in ASCII
			}
			encryptedChar := rune(shift)
			fmt.Printf("Shifted character: %c\n", encryptedChar)
			encryptedString += string(encryptedChar)
		}
	}

	//-------------------------------------------

	return encryptedString
}

// userInputHandler handles user input.
func encryptHandler(c *gin.Context) {
	var data encryptStruct

	// Call BindJSON to bind the received JSON to encStruct.
	if err := c.BindJSON(&data); err != nil {
		fmt.Println("Error binding JSON:", err)
		return
	}

	fmt.Printf("Received text: %s\n", data.Text)
	fmt.Printf("Received key: %d\n", data.Key)
	var encrypted encryptStruct
	encrypted.Text = cesar(data)
	encrypted.Key = data.Key

	fmt.Printf("Encrypted text: %s\n", encrypted.Text)
	// Respond with a JSON containing encStruct.Text and encStruct.Key
	response := gin.H{
		"text": encrypted.Text,
		"key":  encrypted.Key,
	}
	c.JSON(200, response)

}

func main() {
	router := gin.Default()

	router.POST("/encrypt", encryptHandler)

	router.Run(":8080")
}

/*


 */
