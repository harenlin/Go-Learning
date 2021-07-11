package main
import (
	"encoding/base64"
	"fmt"
)

func main() {
	original_text := "I graduated from NCCU this year!"
	encodedText := base64.StdEncoding.EncodeToString([]byte(original_text))
	println("Encoded Text:", encodedText)
	decodedText, err := base64.StdEncoding.DecodeString(encodedText)
	if err == nil {
		fmt.Println("Decoded Text:", string(decodedText))
	} else {
		fmt.Println("Error:", err.Error())
	} 
}

// Encoded Text: SSBncmFkdWF0ZWQgZnJvbSBOQ0NVIHRoaXMgeWVhciE=
// Decoded Text: I graduated from NCCU this year!