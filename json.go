package main
import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	haren := &Person{"Haren Lin", 22}
	json, err := json.Marshal(*haren)
	if err == nil {
		fmt.Println("JSON Format:", string(json))
	} else {
		fmt.Println("Error:", err.Error())
	}
}

// JSON Format: {"Name":"Haren Lin","Age":22}